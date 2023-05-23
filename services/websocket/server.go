/*
 * @Author: psq
 * @Date: 2022-05-08 14:18:08
 * @LastEditors: psq
 * @LastEditTime: 2023-04-29 01:23:12
 */

package websocket

import (
	"gateway-websocket/config"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-module/carbon"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

var (
	upGrader = websocket.Upgrader{
		// 设置消息接收缓冲区大小（byte），如果这个值设置得太小，可能会导致服务端在读取客户端发送的大型消息时遇到问题
		ReadBufferSize: config.GatewayConfig["ReadBufferSize"].(int),
		// 设置消息发送缓冲区大小（byte），如果这个值设置得太小，可能会导致服务端在发送大型消息时遇到问题
		WriteBufferSize: config.GatewayConfig["WriteBufferSize"].(int),
		// 消息包启用压缩
		EnableCompression: config.GatewayConfig["MessageCompression"].(bool),
		// ws握手超时时间
		HandshakeTimeout: time.Duration(config.GatewayConfig["WebsocketHandshakeTimeout"].(int)) * time.Second,
		// ws握手过程中允许跨域
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	// 设置消息包发送格式
	TextMessage = websocket.TextMessage
	// 设置心跳检测间隔时长（秒）
	HeartbeatTime = config.GatewayConfig["HeartbeatTimeout"].(int)
)

func WsServer(c *gin.Context) {

	// 将 HTTP 连接升级为 WebSocket 连接
	conn, err := upGrader.Upgrade(c.Writer, c.Request, nil)

	if err != nil {

		return
	}

	defer conn.Close()

	// 客户端唯一身份标识
	clientID := uuid.New().String()

	// 记录客户端连接信息
	GatewayClients[clientID] = &WebSocketClientBase{
		ID:            clientID,
		Conn:          conn,
		LastHeartbeat: carbon.Now().Timestamp(),
	}

	// 发送客户端唯一标识 ID
	if err = conn.WriteMessage(websocket.TextMessage, []byte(clientID)); err != nil {

		return
	}

	go clientHeartbeatCheck(clientID)

	for {

		// 读取客户端发送过来的消息
		messageType, message, err := conn.ReadMessage()

		// 当收到err时则标识客户端连接出现异常，如断线
		if err != nil {

			if GatewayClients[clientID].BindUid != "" {

				clientUnBindUid(clientID, GatewayClients[clientID].BindUid)
			}

			if len(GatewayClients[clientID].JoinGroup) > 0 {

				clientLeaveGroup(clientID)
			}

			delete(GatewayClients, clientID)
			break
		}

		// 当客户端发送的是心跳时返回pong字符包以此刷新心跳
		if messageType == websocket.TextMessage && string(message) == "ping" {

			err = conn.WriteMessage(websocket.TextMessage, []byte("pong"))

			// 向客户端发送消息如果遇到异常当即断开连接
			if err != nil {

				if GatewayClients[clientID].BindUid != "" {

					clientUnBindUid(clientID, GatewayClients[clientID].BindUid)
				}

				if len(GatewayClients[clientID].JoinGroup) > 0 {

					clientLeaveGroup(clientID)
				}

				delete(GatewayClients, clientID)
				break
			}

			// 刷新最后最后心跳时间
			GatewayClients[clientID].LastHeartbeat = carbon.Now().Timestamp()
		}
	}

}
