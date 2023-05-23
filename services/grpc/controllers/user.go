/*
 * @Author: psq
 * @Date: 2022-08-26 18:19:23
 * @LastEditors: psq
 * @LastEditTime: 2023-05-23 16:05:33
 */
package controllers

import (
	"context"
	"gateway-websocket/services"
	ClientBindUidPB "gateway-websocket/services/grpc/proto/clientbinduid"
	CountOnlineUidPB "gateway-websocket/services/grpc/proto/countonlineuid"
	GetAllOnlineUidPB "gateway-websocket/services/grpc/proto/getallonlineuid"
	GetClientByUidPB "gateway-websocket/services/grpc/proto/getclientbyuid"
	GetUidByClientPB "gateway-websocket/services/grpc/proto/getuidbyclient"
	SendMessageToUidPB "gateway-websocket/services/grpc/proto/sendmessagetouid"
	UidIsOnlinePB "gateway-websocket/services/grpc/proto/uidisonline"
	UnBindUidPB "gateway-websocket/services/grpc/proto/unbinduid"
)

type UserControllers struct{}

var user = new(services.Gateway)

/**
 * @description: 获取Uid下所有client
 * @param {context.Context} ctx
 * @param {*GetUidByClientPB.GetUidByClientRequest} c
 * @return {*}
 */
func (h *UserControllers) GetUidByClient(ctx context.Context, c *GetUidByClientPB.GetUidByClientRequest) (*GetUidByClientPB.GetUidByClientResponse, error) {

	return &GetUidByClientPB.GetUidByClientResponse{Clientid: user.GetUidByClient(c.Uid)}, nil
}

/**
 * @description: 统计在线uid人数
 * @param {context.Context} ctx
 * @param {*CountOnlineUidPB.CountOnlineUidRequest} c
 * @return {*}
 */
func (h *UserControllers) CountOnlineUid(ctx context.Context, c *CountOnlineUidPB.CountOnlineUidRequest) (*CountOnlineUidPB.CountOnlineUidResponse, error) {

	return &CountOnlineUidPB.CountOnlineUidResponse{Count: int32(user.CountOnlineUid())}, nil
}

/**
 * @description: 获取所有在线uid
 * @param {context.Context} ctx
 * @param {*GetAllOnlineUidPB.GetAllOnlineUidRequest} c
 * @return {*}
 */
func (h *UserControllers) GetAllOnlineUid(ctx context.Context, c *GetAllOnlineUidPB.GetAllOnlineUidRequest) (*GetAllOnlineUidPB.GetAllOnlineUidResponse, error) {

	return &GetAllOnlineUidPB.GetAllOnlineUidResponse{Uid: user.GetAllOnlineUid()}, nil
}

/**
 * @description: 向某个uid下绑定的所有client发送消息
 * @param {context.Context} ctx
 * @param {*ClientBindUidPB.ClientBindUidRequest} c
 * @return {*}
 */
func (h *UserControllers) SendMessageToUid(ctx context.Context, c *SendMessageToUidPB.SendMessageToUidRequest) (*SendMessageToUidPB.SendMessageToUidResponse, error) {

	if c.Uid == "" || c.Message == "" {

		return &SendMessageToUidPB.SendMessageToUidResponse{Result: false}, nil
	}

	return &SendMessageToUidPB.SendMessageToUidResponse{Result: user.SendMessageToUid(c.Uid, []byte(c.Message))}, nil
}

/**
 * @description: 判断某个uid是否存在
 * @param {context.Context} ctx
 * @param {*ClientBindUidPB.ClientBindUidRequest} c
 * @return {*}
 */
func (h *UserControllers) UidIsOnline(ctx context.Context, c *UidIsOnlinePB.UidIsOnlineRequest) (*UidIsOnlinePB.UidIsOnlineResponse, error) {

	if c.Uid == "" {

		return &UidIsOnlinePB.UidIsOnlineResponse{Result: false}, nil
	}

	return &UidIsOnlinePB.UidIsOnlineResponse{Result: user.UidIsOnline(c.Uid)}, nil
}

/**
 * @description: client解除绑定Uid
 * @param {context.Context} ctx
 * @param {*ClientBindUidPB.ClientBindUidRequest} c
 * @return {*}
 */
func (h *UserControllers) UnBindUid(ctx context.Context, c *UnBindUidPB.UnBindUidRequest) (*UnBindUidPB.UnBindUidResponse, error) {

	if c.Clientid == "" {

		return &UnBindUidPB.UnBindUidResponse{Result: false}, nil
	}

	return &UnBindUidPB.UnBindUidResponse{Result: user.UnBindUid(c.Clientid)}, nil
}

/**
 * @description: client绑定Uid
 * @param {context.Context} ctx
 * @param {*ClientBindUidPB.ClientBindUidRequest} c
 * @return {*}
 */
func (h *UserControllers) ClientBindUid(ctx context.Context, c *ClientBindUidPB.ClientBindUidRequest) (*ClientBindUidPB.ClientBindUidResponse, error) {

	if c.Clientid == "" || c.Uid == "" {

		return &ClientBindUidPB.ClientBindUidResponse{Result: false}, nil
	}

	return &ClientBindUidPB.ClientBindUidResponse{Result: user.ClientBindUid(c.Clientid, c.Uid)}, nil
}

/**
 * @description: 获取client绑定的Uid
 * @param {context.Context} ctx
 * @param {*GetClientByUidPB.GetClientByUidRequest} c
 * @return {*}
 */
func (h *UserControllers) GetClientByUid(ctx context.Context, c *GetClientByUidPB.GetClientByUidRequest) (*GetClientByUidPB.GetClientByUidResponse, error) {

	if c.Clientid == "" {

		return &GetClientByUidPB.GetClientByUidResponse{Uid: ""}, nil
	}

	return &GetClientByUidPB.GetClientByUidResponse{Uid: user.GetClientByUid(c.Clientid)}, nil
}
