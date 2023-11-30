/*
 * @Author: psq
 * @Date: 2023-04-24 18:47:27
 * @LastEditors: psq
 * @LastEditTime: 2023-11-30 16:13:38
 */

package services

import (
	"gateway-websocket/config"
	"gateway-websocket/services/websocket"
)

type Gateway struct{}

/**
 * @description: 解散某个群组
 * @param {string} groupname
 * @return {*}
 */
func (g Gateway) UnGroup(groupname string) bool {

	groupInterface, ok := websocket.GatewayGroup.Load(groupname)

	if !ok {
		// Group 不存在
		return false
	}

	group, _ := groupInterface.(*websocket.WebSocketGroupBase)

	for _, clientID := range group.ClientID {

		clientInterface, ok := websocket.GatewayClients.Load(clientID)

		if !ok {
			// Client 不存在
			continue
		}

		client, _ := clientInterface.(*websocket.WebSocketClientBase)

		for k, joinedGroup := range client.JoinGroup {
			if joinedGroup == groupname {
				client.JoinGroup = append(client.JoinGroup[:k], client.JoinGroup[k+1:]...)
				break
			}
		}
	}

	websocket.GatewayGroup.Delete(groupname)

	return true
}

/**
 * @description: 发送消息到某个群组
 * @param {string} groupname
 * @param {[]byte} message
 * @return {bool int}
 */
func (g Gateway) SendMessageToGroup(groupname string, message []byte) (int, bool) {

	groupInterface, ok := websocket.GatewayGroup.Load(groupname)

	if !ok {

		return 0, false
	}

	sendCount := 0

	group, _ := groupInterface.(*websocket.WebSocketGroupBase)

	for _, v := range group.ClientID {

		clientInterface, ok := websocket.GatewayClients.Load(v)

		if !ok {

			continue
		}

		client, _ := clientInterface.(*websocket.WebSocketClientBase)

		if err := client.Conn.WriteMessage(config.GatewayConfig["MessageFormat"].(int), message); err == nil {

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

	groupInterface, ok := websocket.GatewayGroup.Load(groupname)

	if !ok {

		return 0
	}

	group, _ := groupInterface.(*websocket.WebSocketGroupBase)

	return len(group.ClientID)
}

/**
 * @description: 统计群组数量
 * @return {int}
 */
func (g Gateway) CountGroup() int {

	length := 0

	websocket.GatewayGroup.Range(func(key, value interface{}) bool {
		// 每次迭代都增加计数
		length++
		return true
	})

	return length
}

/**
 * @description: 获取某个群组内所有在线的client
 * @param {string} groupname
 * @return {[]string}
 */
func (g Gateway) GetGroupOnlineClient(groupname string) []string {

	groupInterface, ok := websocket.GatewayGroup.Load(groupname)

	if !ok {

		return []string{}
	}

	group, _ := groupInterface.(*websocket.WebSocketGroupBase)

	return group.ClientID
}

/**
 * @description: 将client推出群组
 * @param {string} clientid
 * @param {string} groupname
 * @return {bool}
 */
func (g Gateway) LeaveGroup(clientid string, groupname string) bool {

	clientInterface, ok := websocket.GatewayClients.Load(clientid)

	if !ok {

		return false
	}

	groupInterface, ok := websocket.GatewayGroup.Load(groupname)

	if !ok {

		return false
	}

	isLeave := true

	client, _ := clientInterface.(*websocket.WebSocketClientBase)

	// 判断用户是否在需要退出的群组内
	for _, v := range client.JoinGroup {

		if v == groupname {

			isLeave = false
			break
		}
	}

	if isLeave {

		return isLeave
	}

	for i := 0; i < len(client.JoinGroup); i++ {

		if client.JoinGroup[i] == groupname {

			client.JoinGroup = append(client.JoinGroup[:i], client.JoinGroup[i+1:]...)
			break
		}
	}

	group, _ := groupInterface.(*websocket.WebSocketGroupBase)

	for i := 0; i < len(group.ClientID); i++ {

		if group.ClientID[i] == clientid {

			group.ClientID = append(group.ClientID[:i], group.ClientID[i+1:]...)
			break
		}
	}

	if len(group.ClientID) == 0 {

		websocket.GatewayGroup.Delete(groupname)
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

	clientInterface, ok := websocket.GatewayClients.Load(clientid)

	if !ok {

		return false
	}

	client, _ := clientInterface.(*websocket.WebSocketClientBase)

	for _, v := range client.JoinGroup {

		if v == groupname {

			return true
		}
	}

	groupInterface, ok := websocket.GatewayGroup.Load(groupname)

	if !ok {

		websocket.GatewayGroup.Store(groupname, &websocket.WebSocketGroupBase{
			ClientID: []string{clientid},
		})

	} else {

		isExist := false

		group, _ := groupInterface.(*websocket.WebSocketGroupBase)

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

	userInterface, ok := websocket.GatewayUser.Load(uid)

	if !ok {

		return []string{}
	}

	user, _ := userInterface.(*websocket.WebSocketUserBase)

	return user.ClientID
}

/**
 * @description: 统计在线uid数
 * @return {int}
 */
func (g Gateway) CountOnlineUid() int {

	length := 0

	websocket.GatewayUser.Range(func(key, value interface{}) bool {
		// 每次迭代都增加计数
		length++
		return true
	})

	return length
}

/**
 * @description: 获取所有在线的Uid
 * @return {[]string}
 */
func (g Gateway) GetAllOnlineUid() []string {

	var keys []string

	websocket.GatewayUser.Range(func(k, value interface{}) bool {

		user, _ := value.(*websocket.WebSocketUserBase)

		keys = append(keys, user.Uid)

		return true
	})

	return keys
}

/**
 * @description: 发送消息给某个Uid下的所有client
 * @param {string} uid
 * @param {string} message
 * @return {bool}
 */
func (g Gateway) SendMessageToUid(uid string, message []byte) bool {

	userInterface, ok := websocket.GatewayUser.Load(uid)

	if !ok {

		return false
	}

	user, _ := userInterface.(*websocket.WebSocketUserBase)

	for _, v := range user.ClientID {

		if clientInterface, ok := websocket.GatewayClients.Load(v); ok {

			client, _ := clientInterface.(*websocket.WebSocketClientBase)

			client.Conn.WriteMessage(config.GatewayConfig["MessageFormat"].(int), message)

		}
	}

	return true
}

/**
 * @description: 判断某个uid是否在线
 * @param {string} uid
 * @return {bool}
 */
func (g Gateway) UidIsOnline(uid string) bool {

	_, ok := websocket.GatewayUser.Load(uid)

	return ok
}

/**
 * @description: client解除绑定Uid
 * @param {string} clientid
 * @return {bool}
 */
func (g Gateway) UnBindUid(clientid string) bool {

	clientInterface, ok := websocket.GatewayClients.Load(clientid)

	if !ok {

		return false
	}

	client, _ := clientInterface.(*websocket.WebSocketClientBase)

	if client.BindUid == "" {

		return true
	}

	userInterface, ok := websocket.GatewayUser.Load(client.BindUid)

	if !ok {

		return false
	}

	user, _ := userInterface.(*websocket.WebSocketUserBase)

	for k, v := range user.ClientID {

		if v == clientid {

			user.ClientID = append(user.ClientID[:k], user.ClientID[k+1:]...)
		}
	}

	if len(user.ClientID) == 0 {

		websocket.GatewayUser.Delete(client.BindUid)
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

	clientInterface, ok := websocket.GatewayClients.Load(clientid)

	if !ok {

		return false
	}

	client, _ := clientInterface.(*websocket.WebSocketClientBase)

	if client.BindUid != "" {

		g.UnBindUid(clientid)
	}

	client.BindUid = uid

	userInterface, ok := websocket.GatewayUser.Load(uid)

	if !ok {

		websocket.GatewayUser.Store(uid, &websocket.WebSocketUserBase{
			Uid:      uid,
			ClientID: []string{clientid},
		})

	} else {

		isExist := false

		user, _ := userInterface.(*websocket.WebSocketUserBase)

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

	websocket.GatewayClients.Range(func(key, value interface{}) bool {

		client := value.(*websocket.WebSocketClientBase)

		go func(c *websocket.WebSocketClientBase) {

			c.Conn.WriteMessage(config.GatewayConfig["MessageFormat"].(int), message)

		}(client)

		return true
	})
}

/**
 * @description: 发送消息给某个客户端
 * @param {string} clientid
 * @param {string} message
 * @return {bool}
 */
func (g Gateway) SendMessageToClient(clientid string, message []byte) bool {

	clientInterface, ok := websocket.GatewayClients.Load(clientid)

	if !ok {

		return false
	}

	client, _ := clientInterface.(*websocket.WebSocketClientBase)

	return client.Conn.WriteMessage(config.GatewayConfig["MessageFormat"].(int), message) == nil
}

/**
 * @description: 获取client绑定的Uid
 * @param {string} client
 * @return {string}
 */
func (g Gateway) GetClientByUid(clientid string) string {

	clientInterface, ok := websocket.GatewayClients.Load(clientid)

	if !ok {

		return ""
	}

	client, _ := clientInterface.(*websocket.WebSocketClientBase)

	return client.BindUid
}

/**
 * @description: 服务端根据clientid主动关闭客户端连接
 * @param {string} clientid
 * @return {bool}
 */
func (g Gateway) ClonseClient(clientid string) bool {

	clientInterface, ok := websocket.GatewayClients.Load(clientid)

	if !ok {

		return true
	}

	client, _ := clientInterface.(*websocket.WebSocketClientBase)

	client.Conn.Close()
	websocket.GatewayClients.Delete(clientid)

	return true
}

/**
 * @description: 获取所有在线的客户端
 * @return {[]string}
 */
func (g Gateway) GetAllOnlineClient() []string {

	var keys []string

	websocket.GatewayClients.Range(func(k, value interface{}) bool {

		client, _ := value.(*websocket.WebSocketClientBase)

		keys = append(keys, client.ID)

		return true
	})

	return keys
}

/**
 * @description: 判断客户端是否在线（根据clientid）
 * @param {string} clientid
 * @return {bool}
 */
func (g Gateway) ClientIsOnline(clientid string) bool {

	_, ok := websocket.GatewayClients.Load(clientid)

	return ok
}

/**
 * @description: 获取在线客户端数量
 * @return {int}
 */
func (g Gateway) CountOnlineClient() int {

	length := 0

	websocket.GatewayClients.Range(func(key, value interface{}) bool {
		// 每次迭代都增加计数
		length++
		return true
	})

	return length
}
