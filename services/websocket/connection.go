/*
 * @Author: psq
 * @Date: 2023-04-24 15:20:00
 * @LastEditors: psq
 * @LastEditTime: 2024-06-27 15:02:53
 */
package websocket

import (
	"fmt"
	"sync"
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

var GatewayClients, GatewayUser, GatewayGroup sync.Map

/**
 * @description: 客户端心跳检测，超时即断开连接（主要是为了降低服务端承载压力）
 * @param {string} clientID
 * @return {*}
 */
func clientHeartbeatCheck(clientID string) {

	for {

		time.Sleep(5 * time.Second)

		clientInterface, ok := GatewayClients.Load(clientID)

		if !ok {

			break
		}

		client, _ := clientInterface.(*WebSocketClientBase)

		if (carbon.Now().Timestamp() - client.LastHeartbeat) > int64(HeartbeatTime) {

			fmt.Println("Client", clientID, "heartbeat timeout")

			client.Conn.Close()

			if client.BindUid != "" {

				clientUnBindUid(client.ID, client.BindUid)
			}

			clientLeaveGroup(client.ID)

			GatewayClients.Delete(clientID)
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

	value, ok := GatewayUser.Load(uid)

	if ok {

		users := value.(*WebSocketUserBase)

		for k, v := range users.ClientID {

			if v == clientID {

				users.ClientID = append(users.ClientID[:k], users.ClientID[k+1:]...)
			}
		}

		if len(users.ClientID) == 0 {

			GatewayUser.Delete(uid)
		}

	}
}

/**
 * @description: 客户端断线时自动踢出已加入的群组
 * @param {string} clientID
 * @return {*}
 */
func clientLeaveGroup(clientID string) {
	// 使用 Load 方法获取值
	value, ok := GatewayClients.Load(clientID)
	if !ok {
		// 如果没有找到对应的值，处理相应的逻辑
		return
	}

	client := value.(*WebSocketClientBase)

	// 遍历 JoinGroup
	for _, v := range client.JoinGroup {
		// 使用 Load 方法获取值
		groupValue, groupOK := GatewayGroup.Load(v)
		if !groupOK {
			// 如果没有找到对应的值，处理相应的逻辑
			continue
		}

		group := groupValue.(*WebSocketGroupBase)

		// 在群组中找到对应的 clientID，并删除
		for j, id := range group.ClientID {
			if id == clientID {
				copy(group.ClientID[j:], group.ClientID[j+1:])
				group.ClientID = group.ClientID[:len(group.ClientID)-1]

				// 如果群组中没有成员了，删除群组
				if len(group.ClientID) == 0 {
					GatewayGroup.Delete(v)
				}

				break
			}
		}
	}
}
