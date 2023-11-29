/*
 * @Author: psq
 * @Date: 2023-04-24 15:20:00
 * @LastEditors: psq
 * @LastEditTime: 2023-08-04 16:21:26
 */
package websocket

import (
	"fmt"
	"time"

	"github.com/golang-module/carbon"
	"github.com/gorilla/websocket"
)

type WebSocketClientBase struct {
	ID            string
	Conn          *websocket.Conn
	LastHeartbeat int64
	BindUid       string
	JoinGroup     []string
}

type WebSocketUserBase struct {
	Uid      string
	ClientID []string
}

type WebSocketGroupBase struct {
	ClientID []string
}

var GatewayClients = make(map[string]*WebSocketClientBase)

var GatewayUser = make(map[string]*WebSocketUserBase)

var GatewayGroup = make(map[string]*WebSocketGroupBase)

/**
 * @description: 客户端心跳检测，超时即断开连接（主要是为了降低服务端承载压力）
 * @param {string} clientID
 * @return {*}
 */
func clientHeartbeatCheck(clientID string) {

	for {

		time.Sleep(5 * time.Second)

		client, exists := GatewayClients[clientID]

		if !exists {

			break
		}

		if (carbon.Now().Timestamp() - client.LastHeartbeat) > int64(HeartbeatTime) {

			fmt.Println("Client", clientID, "heartbeat timeout")

			client.Conn.Close()
			delete(GatewayClients, clientID)
			break
		}
	}
}

/**
 * @description: 客户端断线时自动踢出Uid绑定列表
 * @param {string} clientID
 * @param {string} uid
 * @return {*}
 */
func clientUnBindUid(clientID string, uid string) {

	for k, v := range GatewayUser[uid].ClientID {

		if v == clientID {

			GatewayUser[uid].ClientID = append(GatewayUser[uid].ClientID[:k], GatewayUser[uid].ClientID[k+1:]...)
		}
	}

	if len(GatewayUser[uid].ClientID) == 0 {

		delete(GatewayUser, uid)
	}
}

/**
 * @description: 客户端断线时自动踢出已加入的群组
 * @param {string} clientID
 * @return {*}
 */
func clientLeaveGroup(clientID string) {

	for _, v := range GatewayClients[clientID].JoinGroup {

		for j, id := range GatewayGroup[v].ClientID {

			if id == clientID {

				copy(GatewayGroup[v].ClientID[j:], GatewayGroup[v].ClientID[j+1:])
				GatewayGroup[v].ClientID = GatewayGroup[v].ClientID[:len(GatewayGroup[v].ClientID)-1]

				if len(GatewayGroup[v].ClientID) == 0 {

					delete(GatewayGroup, v)
				}

				break
			}
		}
	}

}
