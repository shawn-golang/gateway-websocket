/*
 * @Author: psq
 * @Date: 2022-08-26 18:19:23
 * @LastEditors: psq
 * @LastEditTime: 2023-05-06 15:53:56
 */
package controllers

import (
	"context"
	"gateway-websocket/services"

	CountGroupPB "gateway-websocket/services/grpc/proto/countgroup"
	CountOnlineGroupPB "gateway-websocket/services/grpc/proto/countonlinegroup"
	GetGroupOnlineClientPB "gateway-websocket/services/grpc/proto/getgrouponlineclient"
	JoinGroupPB "gateway-websocket/services/grpc/proto/joingroup"
	LeaveGroupPB "gateway-websocket/services/grpc/proto/leavegroup"
	SendMessageToGroupPB "gateway-websocket/services/grpc/proto/sendmessagetogroup"
	UnGroupPB "gateway-websocket/services/grpc/proto/ungroup"
)

type GroupControllers struct{}

var group = new(services.Gateway)

/**
 * @description: 解散某个群组
 * @param {context.Context} ctx
 * @param {*UnGroupPB.UnGroupRequest} c
 * @return {*}
 */
func (h *GroupControllers) UnGroup(ctx context.Context, c *UnGroupPB.UnGroupRequest) (*UnGroupPB.UnGroupResponse, error) {

	return &UnGroupPB.UnGroupResponse{Result: group.UnGroup(c.Group)}, nil
}

/**
 * @description: 向某个群组发送消息
 * @param {context.Context} ctx
 * @param {*SendMessageToGroupPB.SendMessageToGroupRequest} c
 * @return {*}
 */
func (h *GroupControllers) SendMessageToGroup(ctx context.Context, c *SendMessageToGroupPB.SendMessageToGroupRequest) (*SendMessageToGroupPB.SendMessageToGroupResponse, error) {

	arrive, ok := group.SendMessageToGroup(c.Group, []byte(c.Message))

	return &SendMessageToGroupPB.SendMessageToGroupResponse{Result: ok, Arrive: int32(arrive)}, nil
}

/**
 * @description: 统计群组内在线的client数量
 * @param {context.Context} ctx
 * @param {*CountOnlineGroupPB.CountOnlineGroupRequest} c
 * @return {*}
 */
func (h *GroupControllers) CountOnlineGroup(ctx context.Context, c *CountOnlineGroupPB.CountOnlineGroupRequest) (*CountOnlineGroupPB.CountOnlineGroupResponse, error) {

	return &CountOnlineGroupPB.CountOnlineGroupResponse{Count: int32(group.CountOnlineGroup(c.Group))}, nil
}

/**
 * @description: 统计群组数量
 * @param {context.Context} ctx
 * @param {*GetGroupOnlineClientPB.GetGroupOnlineClientRequest} c
 * @return {*}
 */
func (h *GroupControllers) CountGroup(ctx context.Context, c *CountGroupPB.CountGroupRequest) (*CountGroupPB.CountGroupResponse, error) {

	return &CountGroupPB.CountGroupResponse{Count: int32(group.CountGroup())}, nil
}

/**
 * @description: 获取群组内所有在线的client
 * @param {context.Context} ctx
 * @param {*GetGroupOnlineClientPB.GetGroupOnlineClientRequest} c
 * @return {*}
 */
func (h *GroupControllers) GetGroupOnlineClient(ctx context.Context, c *GetGroupOnlineClientPB.GetGroupOnlineClientRequest) (*GetGroupOnlineClientPB.GetGroupOnlineClientResponse, error) {

	return &GetGroupOnlineClientPB.GetGroupOnlineClientResponse{Clientid: group.GetGroupOnlineClient(c.Group)}, nil
}

/**
 * @description: 将client移出群组
 * @param {context.Context} ctx
 * @param {*LeaveGroupPB.LeaveGroupRequest} c
 * @return {*}
 */
func (h *GroupControllers) LeaveGroup(ctx context.Context, c *LeaveGroupPB.LeaveGroupRequest) (*LeaveGroupPB.LeaveGroupResponse, error) {

	if c.Clientid == "" || c.Group == "" {

		return &LeaveGroupPB.LeaveGroupResponse{Result: false}, nil
	}

	return &LeaveGroupPB.LeaveGroupResponse{Result: group.LeaveGroup(c.Clientid, c.Group)}, nil
}

/**
 * @description: 将client加入群组
 * @param {context.Context} ctx
 * @param {*JoinGroupPB.JoinGroupRequest} c
 * @return {*}
 */
func (h *GroupControllers) JoinGroup(ctx context.Context, c *JoinGroupPB.JoinGroupRequest) (*JoinGroupPB.JoinGroupResponse, error) {

	if c.Clientid == "" || c.Group == "" {

		return &JoinGroupPB.JoinGroupResponse{Result: false}, nil
	}

	return &JoinGroupPB.JoinGroupResponse{Result: group.JoinGroup(c.Clientid, c.Group)}, nil
}
