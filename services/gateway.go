/*
 * @Author: psq
 * @Date: 2023-04-24 18:47:27
 * @LastEditors: psq
 * @LastEditTime: 2023-05-06 10:53:35
 */

package services

import (
	"fmt"
	"gateway-websocket/services/websocket"
)

type Gateway struct{}

/**
 * @description: 解散某个群组
 * @param {string} groupname
 * @return {*}
 */
func (g Gateway) UnGroup(groupname string) bool {

	group, ok := websocket.GatewayGroup[groupname]

	if !ok {

		return true
	}

	fmt.Println(group.ClientID)

	for _, v := range group.ClientID {

		client, ok := websocket.GatewayClients[v]

		if !ok {

			continue
		}

		for k, c := range client.JoinGroup {

			if c == groupname {
				client.JoinGroup = append(client.JoinGroup[:k], client.JoinGroup[k+1:]...)
				break
			}
		}

	}

	delete(websocket.GatewayGroup, groupname)

	return true
}

/**
 * @description: 发送消息到某个群组
 * @param {string} groupname
 * @param {[]byte} message
 * @return {bool int}
 */
func (g Gateway) SendMessageToGroup(groupname string, message []byte) (int, bool) {

	group, ok := websocket.GatewayGroup[groupname]

	if !ok {

		return 0, false
	}

	sendCount := 0

	for _, v := range group.ClientID {

		client, ok := websocket.GatewayClients[v]

		if !ok {

			continue
		}

		if err := client.Conn.WriteMessage(websocket.TextMessage, message); err == nil {

			sendCount++
		}
	}

	return sendCount, true
}

/**
 * @description: 统计某个群组在线成员数
 * @param {string} groupname
 * @return {int}
 */
func (g Gateway) CountOnlineGroup(groupname string) int {

	group, ok := websocket.GatewayGroup[groupname]

	if !ok {

		return 0
	}

	return len(group.ClientID)
}

/**
 * @description: 统计群组数量
 * @return {int}
 */
func (g Gateway) CountGroup() int {

	return len(websocket.GatewayGroup)
}

/**
 * @description: 获取某个群组内所有在线的client
 * @param {string} groupname
 * @return {[]string}
 */
func (g Gateway) GetGroupOnlineClient(groupname string) []string {

	group, ok := websocket.GatewayGroup[groupname]

	if !ok {

		return []string{}
	}

	return group.ClientID
}

/**
 * @description: 将client推出群组
 * @param {string} clientid
 * @param {string} groupname
 * @return {bool}
 */
func (g Gateway) LeaveGroup(clientid string, groupname string) bool {

	client, ok := websocket.GatewayClients[clientid]

	if !ok {

		return false
	}

	isLeave := true

	for _, v := range client.JoinGroup {

		if v == groupname {

			isLeave = false
			break
		}
	}

	if isLeave {

		return isLeave
	}

	for i := 0; i < len(websocket.GatewayClients[clientid].JoinGroup); i++ {

		if websocket.GatewayClients[clientid].JoinGroup[i] == groupname {

			websocket.GatewayClients[clientid].JoinGroup = append(websocket.GatewayClients[clientid].JoinGroup[:i], websocket.GatewayClients[clientid].JoinGroup[i+1:]...)
			break
		}
	}

	for i := 0; i < len(websocket.GatewayGroup[groupname].ClientID); i++ {

		if websocket.GatewayGroup[groupname].ClientID[i] == clientid {

			websocket.GatewayGroup[groupname].ClientID = append(websocket.GatewayGroup[groupname].ClientID[:i], websocket.GatewayGroup[groupname].ClientID[i+1:]...)
			break
		}
	}

	if len(websocket.GatewayGroup[groupname].ClientID) == 0 {

		delete(websocket.GatewayGroup, groupname)
	}

	return true
}

/**
 * @description: 将client加入群组
 * @param {string} clientid
 * @param {string} group
 * @return {bool}
 */
func (g Gateway) JoinGroup(clientid string, groupname string) bool {

	client, ok := websocket.GatewayClients[clientid]

	if !ok {

		return false
	}

	for _, v := range client.JoinGroup {

		if v == groupname {

			return true
		}
	}

	group, ok := websocket.GatewayGroup[groupname]

	if !ok {

		websocket.GatewayGroup[groupname] = &websocket.WebSocketGroupBase{
			ClientID: []string{clientid},
		}

	} else {

		isExist := false

		for _, v := range group.ClientID {

			if v == clientid {

				isExist = true
				break
			}
		}

		if !isExist {

			group.ClientID = append(group.ClientID, clientid)
		}
	}

	client.JoinGroup = append(client.JoinGroup, groupname)

	return true
}

/**
 * @description: 获取Uid下所有client
 * @param {string} uid
 * @return {[]string}
 */
func (g Gateway) GetUidByClient(uid string) []string {

	user, ok := websocket.GatewayUser[uid]

	if !ok {

		return []string{}
	}

	return user.ClientID
}

/**
 * @description: 统计在线uid数
 * @return {int}
 */
func (g Gateway) CountOnlineUid() int {

	return len(websocket.GatewayUser)
}

/**
 * @description: 获取所有在线的Uid
 * @return {[]string}
 */
func (g Gateway) GetAllOnlineUid() []string {

	keys := make([]string, len(websocket.GatewayUser))

	i := 0

	for k := range websocket.GatewayUser {

		keys[i] = k
		i++
	}

	return keys
}

/**
 * @description: 发送消息给某个Uid下的所有client
 * @param {string} uid
 * @param {string} message
 * @return {bool}
 */
func (g Gateway) SendMessageToUid(uid string, message []byte) bool {

	user, ok := websocket.GatewayUser[uid]

	if !ok {

		return false
	}

	for _, v := range user.ClientID {

		_ = websocket.GatewayClients[v].Conn.WriteMessage(websocket.TextMessage, message)
	}

	return true
}

/**
 * @description: 判断某个uid是否在线
 * @param {string} uid
 * @return {bool}
 */
func (g Gateway) UidIsOnline(uid string) bool {

	_, ok := websocket.GatewayUser[uid]

	return ok
}

/**
 * @description: client解除绑定Uid
 * @param {string} clientid
 * @return {bool}
 */
func (g Gateway) UnBindUid(clientid string) bool {

	client, ok := websocket.GatewayClients[clientid]

	if !ok {

		return false
	}

	for k, v := range websocket.GatewayUser[client.BindUid].ClientID {

		if v == clientid {

			websocket.GatewayUser[client.BindUid].ClientID = append(websocket.GatewayUser[client.BindUid].ClientID[:k], websocket.GatewayUser[client.BindUid].ClientID[k+1:]...)
		}
	}

	if len(websocket.GatewayUser[client.BindUid].ClientID) == 0 {

		delete(websocket.GatewayUser, client.BindUid)
	}

	client.BindUid = ""

	return client.BindUid == ""
}

/**
 * @description: client绑定Uid（一个Uid可以绑定多个client，但一个client只能绑定一个uid）
 * @param {string} clientid
 * @param {string} uid
 * @return {bool}
 */
func (g Gateway) ClientBindUid(clientid string, uid string) bool {

	client, ok := websocket.GatewayClients[clientid]

	if !ok {

		return false
	}

	client.BindUid = uid

	user, ok := websocket.GatewayUser[uid]

	if !ok {

		websocket.GatewayUser[uid] = &websocket.WebSocketUserBase{
			Uid:      uid,
			ClientID: []string{clientid},
		}

	} else {

		isExist := false

		for _, v := range user.ClientID {

			if v == clientid {

				isExist = true
				break
			}
		}

		if !isExist {

			user.ClientID = append(user.ClientID, clientid)
		}
	}

	return client.BindUid == uid
}

/**
 * @description: 消息广播（向所有客户端发送一条同样的消息，这里需要注意，广播时不要发送太大的消息包）
 * @param {[]byte} message
 * @return {*}
 */
func (g Gateway) BroadcastMessage(message []byte) {

	for _, client := range websocket.GatewayClients {

		go func(c *websocket.WebSocketClientBase) {

			_ = c.Conn.WriteMessage(websocket.TextMessage, message)

		}(client)
	}
}

/**
 * @description: 发送消息给某个客户端
 * @param {string} clientid
 * @param {string} message
 * @return {bool}
 */
func (g Gateway) SendMessageToClient(clientid string, message []byte) bool {

	client, ok := websocket.GatewayClients[clientid]

	if !ok {

		return false
	}

	if err := client.Conn.WriteMessage(websocket.TextMessage, message); err != nil {

		return false
	}

	return true
}

/**
 * @description: 获取client绑定的Uid
 * @param {string} client
 * @return {string}
 */
func (g Gateway) GetClientByUid(clientid string) string {

	client, ok := websocket.GatewayClients[clientid]

	if !ok {

		return ""
	}

	return client.BindUid
}

/**
 * @description: 服务端根据clientid主动关闭客户端连接
 * @param {string} clientid
 * @return {bool}
 */
func (g Gateway) ClonseClient(clientid string) bool {

	client, exists := websocket.GatewayClients[clientid]

	if !exists {

		return false
	}

	client.Conn.Close()
	delete(websocket.GatewayClients, clientid)

	return true
}

/**
 * @description: 获取所有在线的客户端
 * @return {[]string}
 */
func (g Gateway) GetAllOnlineClient() []string {

	keys := make([]string, len(websocket.GatewayClients))

	i := 0

	for k := range websocket.GatewayClients {

		keys[i] = k
		i++
	}

	return keys
}

/**
 * @description: 判断客户端是否在线（根据clientid）
 * @param {string} clientid
 * @return {bool}
 */
func (g Gateway) ClientIsOnline(clientid string) bool {

	_, ok := websocket.GatewayClients[clientid]

	return ok
}

/**
 * @description: 获取在线客户端数量
 * @return {int}
 */
func (g Gateway) CountOnlineClient() int {

	return len(websocket.GatewayClients)
}
