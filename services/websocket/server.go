/*
 * @Author: psq
 * @Date: 2022-05-08 14:18:08
 * @LastEditors: psq
 * @LastEditTime: 2024-07-24 11:51:25
 */

package websocket

import (
	"fmt"
	"gateway-websocket/config"
	"net/http"
	"time"

	"github.com/golang-module/carbon"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

var (
	socketSet = websocket.Upgrader{
		// 设置消息发送缓冲区大小（byte），如果这个值设置得太小，可能会导致服务端在发送大型消息时遇到问题
		WriteBufferSize: config.GatewayConfig["WriteBufferSize"].(int),
		// 消息包启用压缩
		EnableCompression: true,
		// ws握手超时时间
		HandshakeTimeout: 5 * time.Second,
		// ws握手过程中允许跨域
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	// 设置心跳检测间隔时长（秒）
	HeartbeatTime = config.GatewayConfig["HeartbeatTimeout"].(int)
)

/**
 * @description: 初始化客户端连接
 * @param {*websocket.Conn} conn
 * @return {*}
 */
func handleClientInit(conn *websocket.Conn) string {

	// clientID也可以用其他方式生成，只要能保证在所有服务端中都能保证唯一即可
	clientID := uuid.New().String()

	// 使用 Store 方法存储值
	GatewayClients.Store(clientID, &WebSocketClientBase{
		ID:            clientID,
		Conn:          conn,
		LastHeartbeat: carbon.Now().Timestamp(),
	})

	if err := conn.WriteMessage(config.GatewayConfig["MessageFormat"].(int), []byte(clientID)); err != nil {

		handleClientDisconnect(clientID)
		return ""
	}

	return clientID
}

/**
 * @description: 主动关闭客户端连接
 * @param {string} clientID
 * @return {*}
 */
func handleClientDisconnect(clientID string) {

	// 使用 Load 和 Delete 方法，不需要额外的锁定操作
	if v, ok := GatewayClients.Load(clientID); ok {

		client := v.(*WebSocketClientBase)

		if client.BindUid != "" {
			clientUnBindUid(clientID, client.BindUid)
		}

		if len(client.JoinGroup) > 0 {
			clientLeaveGroup(clientID)
		}

		GatewayClients.Delete(clientID)
	}
}

/**
 * @description: 向客户端回复心跳消息
 * @param {*websocket.Conn} conn
 * @param {string} clientID
 * @param {int} messageType
 * @param {[]byte} message
 * @return {*}
 */
func handleClientMessage(conn *websocket.Conn, clientID string, messageType int, message []byte) {

	// 使用 Load 方法获取值
	v, ok := GatewayClients.Load(clientID)
	if !ok {
		// 如果没有找到对应的值，处理相应的逻辑
		handleClientDisconnect(clientID)
		return
	}

	client := v.(*WebSocketClientBase)

	if messageType == config.GatewayConfig["MessageFormat"].(int) && string(message) == "ping" {

		if err := conn.WriteMessage(config.GatewayConfig["MessageFormat"].(int), []byte("pong")); err != nil {

			handleClientDisconnect(clientID)
			return
		}

		GatewayClients.Store(clientID, &WebSocketClientBase{
			ID:            clientID,
			Conn:          conn,
			LastHeartbeat: carbon.Now().Timestamp(),
			BindUid:       client.BindUid,
			JoinGroup:     client.JoinGroup,
		})
	}
}

/**
 * @description: 使用net/http包创建服务根路由
 * @param {*gin.Context} c
 * @return {*}
 */
//func WsServer(c *gin.Context) {
func WsServer(w http.ResponseWriter, r *http.Request) {

	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("WsServer panic: %v\n", err)
		}
	}()

	// 使用gin包将请求升级为 WebSocket 连接
	//conn, err := socketSet.Upgrade(c.Writer, c.Request, nil)

	// 使用net/http包将请求升级为 WebSocket 连接
	conn, err := socketSet.Upgrade(w, r, nil)

	if err != nil {
		return
	}

	defer conn.Close()

	// 客户端唯一身份标识
	clientID := handleClientInit(conn)

	// 发送客户端唯一标识 ID
	if clientID == "" {
		return
	}

	// 创建一个独立的协程监听这个客户端连接，当超过一定时间没有收到心跳请求时服务端主动断开连接
	go clientHeartbeatCheck(clientID)

	for {

		// 读取客户端发送过来的消息
		messageType, message, err := conn.ReadMessage()

		// 当收到err时则标识客户端连接出现异常，如断线
		if err != nil {

			handleClientDisconnect(clientID)
			continue
		}

		handleClientMessage(conn, clientID, messageType, message)
	}

}
