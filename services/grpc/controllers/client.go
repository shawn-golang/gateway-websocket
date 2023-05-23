/*
 * @Author: psq
 * @Date: 2022-08-26 18:19:23
 * @LastEditors: psq
 * @LastEditTime: 2023-05-06 14:48:46
 */
package controllers

import (
	"context"
	"gateway-websocket/services"
	BroadcastMessagePB "gateway-websocket/services/grpc/proto/broadcastmessage"
	ClientIsOnlinePB "gateway-websocket/services/grpc/proto/clientisonline"
	ClonseClientPB "gateway-websocket/services/grpc/proto/clonseclient"
	CountOnlineClientPB "gateway-websocket/services/grpc/proto/countonlineclient"
	GetAllOnlineClientPB "gateway-websocket/services/grpc/proto/getallonlineclient"
	SendMessageToClientPB "gateway-websocket/services/grpc/proto/sendmessagetoclient"
)

type ClientControllers struct{}

var client = new(services.Gateway)

/**
 * @description: 向所有client广播消息
 * @param {context.Context} ctx
 * @param {*SendMessageToClientPB.SendMessageToClientRequest} c
 * @return {*}
 */
func (h *ClientControllers) BroadcastMessage(ctx context.Context, c *BroadcastMessagePB.BroadcastMessageRequest) (*BroadcastMessagePB.BroadcastMessageResponse, error) {

	if c.Message == "" {

		return &BroadcastMessagePB.BroadcastMessageResponse{Result: false}, nil
	}

	client.BroadcastMessage([]byte(c.Message))

	return &BroadcastMessagePB.BroadcastMessageResponse{Result: true}, nil
}

/**
 * @description: 向某个client发送消息
 * @param {context.Context} ctx
 * @param {*ClientBindUidPB.ClientBindUidRequest} c
 * @return {*}
 */
func (h *ClientControllers) SendMessageToClient(ctx context.Context, c *SendMessageToClientPB.SendMessageToClientRequest) (*SendMessageToClientPB.SendMessageToClientResponse, error) {

	if c.Clientid == "" || c.Message == "" {

		return &SendMessageToClientPB.SendMessageToClientResponse{Result: false}, nil
	}

	return &SendMessageToClientPB.SendMessageToClientResponse{Result: client.SendMessageToClient(c.Clientid, []byte(c.Message))}, nil
}

/**
 * @description: 关闭某个client的连接
 * @param {context.Context} ctx
 * @param {*ClonseClientPB.ClonseClientRequest} c
 * @return {*}
 */
func (h *ClientControllers) ClonseClient(ctx context.Context, c *ClonseClientPB.ClonseClientRequest) (*ClonseClientPB.ClonseClientResponse, error) {

	if c.Clientid == "" {

		return &ClonseClientPB.ClonseClientResponse{Result: false}, nil
	}

	return &ClonseClientPB.ClonseClientResponse{Result: client.ClonseClient(c.Clientid)}, nil
}

/**
 * @description: 获取所有在线客户端
 * @param {context.Context} ctx
 * @param {*GetAllOnlineClientPB.GetAllOnlineClientRequest} c
 * @return {*}
 */
func (h *ClientControllers) GetAllOnlineClient(ctx context.Context, c *GetAllOnlineClientPB.GetAllOnlineClientRequest) (*GetAllOnlineClientPB.GetAllOnlineClientResponse, error) {

	return &GetAllOnlineClientPB.GetAllOnlineClientResponse{Client: client.GetAllOnlineClient()}, nil
}

/**
 * @description: 判断client是否在线
 * @param {context.Context} ctx
 * @param {*ClientIsOnlinePB.ClientIsOnlineRequest} c
 * @return {*}
 */
func (h *ClientControllers) ClientIsOnline(ctx context.Context, c *ClientIsOnlinePB.ClientIsOnlineRequest) (*ClientIsOnlinePB.ClientIsOnlineResponse, error) {

	if c.Clientid == "" {

		return &ClientIsOnlinePB.ClientIsOnlineResponse{Isonline: false}, nil
	}

	return &ClientIsOnlinePB.ClientIsOnlineResponse{Isonline: client.ClientIsOnline(c.Clientid)}, nil
}

/**
 * @description: 统计在线client数量
 * @param {context.Context} ctx
 * @param {*CountOnlineClientPB.CountOnlineClientRequest} c
 * @return {*}
 */
func (h *ClientControllers) CountOnlineClient(ctx context.Context, c *CountOnlineClientPB.CountOnlineClientRequest) (*CountOnlineClientPB.CountOnlineClientResponse, error) {

	return &CountOnlineClientPB.CountOnlineClientResponse{Count: int32(client.CountOnlineClient())}, nil
}
