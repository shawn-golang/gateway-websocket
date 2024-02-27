/*
 * @Author: psq
 * @Date: 2024-02-26 17:03:01
 * @LastEditors: psq
 * @LastEditTime: 2024-02-27 11:33:16
 */
package services

import (
	"context"
	"fmt"
	BroadcastMessagePB "gateway-websocket/services/grpc/proto/broadcastmessage"
	ClientBindUidPB "gateway-websocket/services/grpc/proto/clientbinduid"
	ClientIsOnlinePB "gateway-websocket/services/grpc/proto/clientisonline"
	ClonseClientPB "gateway-websocket/services/grpc/proto/clonseclient"
	CountOnlineClientPB "gateway-websocket/services/grpc/proto/countonlineclient"
	GetAllOnlineClientPB "gateway-websocket/services/grpc/proto/getallonlineclient"
	GetClientByUidPB "gateway-websocket/services/grpc/proto/getclientbyuid"
	SendMessageToClientPB "gateway-websocket/services/grpc/proto/sendmessagetoclient"

	CountOnlineUidPB "gateway-websocket/services/grpc/proto/countonlineuid"
	GetAllOnlineUidPB "gateway-websocket/services/grpc/proto/getallonlineuid"
	GetUidByClientPB "gateway-websocket/services/grpc/proto/getuidbyclient"
	SendMessageToUidPB "gateway-websocket/services/grpc/proto/sendmessagetouid"
	UidIsOnlinePB "gateway-websocket/services/grpc/proto/uidisonline"
	UnBindUidPB "gateway-websocket/services/grpc/proto/unbinduid"

	CountGroupPB "gateway-websocket/services/grpc/proto/countgroup"
	CountOnlineGroupPB "gateway-websocket/services/grpc/proto/countonlinegroup"
	GetGroupOnlineClientPB "gateway-websocket/services/grpc/proto/getgrouponlineclient"
	JoinGroupPB "gateway-websocket/services/grpc/proto/joingroup"
	LeaveGroupPB "gateway-websocket/services/grpc/proto/leavegroup"
	SendMessageToGroupPB "gateway-websocket/services/grpc/proto/sendmessagetogroup"
	UnGroupPB "gateway-websocket/services/grpc/proto/ungroup"

	"runtime"

	"google.golang.org/grpc"
)

type GatewayWebSocketClient struct {
	conn *grpc.ClientConn
}

/**
 * @description: 解散某个群组
 * @param {string} group
 * @return {*}
 */
func (c *GatewayWebSocketClient) UnGroup(group string) (bool, error) {

	Request := &UnGroupPB.UnGroupRequest{}
	Request.Group = group

	Response, err := UnGroupPB.NewUnGroupClient(c.conn).UnGroup(context.Background(), Request)

	if err != nil {

		return false, err
	}

	return Response.Result, nil
}

/**
 * @description: 向某个群组发送消息
 * @param {string} group
 * @param {string} message
 * @return {*}
 */
func (c *GatewayWebSocketClient) SendMessageToGroup(group string, message string) (bool, error) {

	Request := &SendMessageToGroupPB.SendMessageToGroupRequest{}
	Request.Group = group
	Request.Message = message

	Response, err := SendMessageToGroupPB.NewSendMessageToGroupClient(c.conn).SendMessageToGroup(context.Background(), Request)

	if err != nil {

		return false, err
	}

	return Response.Result, nil
}

/**
 * @description: 将client移出群组
 * @param {string} clientid
 * @param {string} group
 * @return {*}
 */
func (c *GatewayWebSocketClient) LeaveGroup(clientid string, group string) (bool, error) {

	Request := &LeaveGroupPB.LeaveGroupRequest{}
	Request.Clientid = clientid
	Request.Group = group

	Response, err := LeaveGroupPB.NewLeaveGroupClient(c.conn).LeaveGroup(context.Background(), Request)

	if err != nil {

		return false, err
	}

	return Response.Result, nil
}

/**
 * @description: 将client加入群组
 * @param {string} clientid
 * @param {string} group
 * @return {*}
 */
func (c *GatewayWebSocketClient) JoinGroup(clientid string, group string) (bool, error) {

	Request := &JoinGroupPB.JoinGroupRequest{}
	Request.Clientid = clientid
	Request.Group = group

	Response, err := JoinGroupPB.NewJoinGroupClient(c.conn).JoinGroup(context.Background(), Request)

	if err != nil {

		return false, err
	}

	return Response.Result, nil
}

/**
 * @description: 获取群组内所有在线的client
 * @param {string} group
 * @return {*}
 */
func (c *GatewayWebSocketClient) GetGroupOnlineClient(group string) ([]string, error) {

	Request := &GetGroupOnlineClientPB.GetGroupOnlineClientRequest{}
	Request.Group = group

	Response, err := GetGroupOnlineClientPB.NewGetGroupOnlineClientClient(c.conn).GetGroupOnlineClient(context.Background(), Request)

	if err != nil {

		return []string{}, err
	}

	return Response.Clientid, nil
}

/**
 * @description: 统计群组内在线的client数量
 * @param {string} group
 * @return {*}
 */
func (c *GatewayWebSocketClient) CountOnlineGroup(group string) (int32, error) {

	Request := &CountOnlineGroupPB.CountOnlineGroupRequest{}
	Request.Group = group

	Response, err := CountOnlineGroupPB.NewCountOnlineGroupClient(c.conn).CountOnlineGroup(context.Background(), Request)

	if err != nil {

		return 0, err
	}

	return Response.Count, nil
}

/**
 * @description: 统计群组数量
 * @return {*}
 */
func (c *GatewayWebSocketClient) CountGroup() (int32, error) {

	Request := &CountGroupPB.CountGroupRequest{}

	Response, err := CountGroupPB.NewCountGroupClient(c.conn).CountGroup(context.Background(), Request)

	if err != nil {

		return 0, err
	}

	return Response.Count, nil
}

/**
 * @description: client解除绑定Uid
 * @param {string} clientid
 * @return {*}
 */
func (c *GatewayWebSocketClient) UnBindUid(clientid string) (bool, error) {

	Request := &UnBindUidPB.UnBindUidRequest{}
	Request.Clientid = clientid

	Response, err := UnBindUidPB.NewUnBindUidClient(c.conn).UnBindUid(context.Background(), Request)

	if err != nil {

		return false, err
	}

	return Response.Result, nil
}

/**
 * @description: 判断某个uid是否存在
 * @param {string} uid
 * @return {*}
 */
func (c *GatewayWebSocketClient) UidIsOnline(uid string) (bool, error) {

	Request := &UidIsOnlinePB.UidIsOnlineRequest{}
	Request.Uid = uid

	Response, err := UidIsOnlinePB.NewUidIsOnlineClient(c.conn).UidIsOnline(context.Background(), Request)

	if err != nil {

		return false, err
	}

	return Response.Result, nil
}

/**
 * @description: 向某个uid下绑定的所有client发送消息
 * @param {string} uid
 * @param {string} message
 * @return {*}
 */
func (c *GatewayWebSocketClient) SendMessageToUid(uid string, message string) (bool, error) {

	Request := &SendMessageToUidPB.SendMessageToUidRequest{}
	Request.Uid = uid
	Request.Message = message

	Response, err := SendMessageToUidPB.NewSendMessageToUidClient(c.conn).SendMessageToUid(context.Background(), Request)

	if err != nil {

		return false, err
	}

	return Response.Result, nil
}

/**
 * @description: 获取Uid下所有client
 * @param {string} uid
 * @return {*}
 */
func (c *GatewayWebSocketClient) GetUidByClient(uid string) ([]string, error) {

	Request := &GetUidByClientPB.GetUidByClientRequest{}
	Request.Uid = uid

	Response, err := GetUidByClientPB.NewGetUidByClientClient(c.conn).GetUidByClient(context.Background(), Request)

	if err != nil {

		return []string{}, err
	}

	return Response.Clientid, nil
}

/**
 * @description: 获取所有在线uid
 * @return {*}
 */
func (c *GatewayWebSocketClient) GetAllOnlineUid() ([]string, error) {

	Request := &GetAllOnlineUidPB.GetAllOnlineUidRequest{}

	Response, err := GetAllOnlineUidPB.NewGetAllOnlineUidClient(c.conn).GetAllOnlineUid(context.Background(), Request)

	if err != nil {

		return []string{}, err
	}

	return Response.Uid, nil
}

/**
 * @description: 统计在线uid人数
 * @return {*}
 */
func (c *GatewayWebSocketClient) CountOnlineUid() (int32, error) {

	Request := &CountOnlineUidPB.CountOnlineUidRequest{}

	Response, err := CountOnlineUidPB.NewCountOnlineUidClient(c.conn).CountOnlineUid(context.Background(), Request)

	if err != nil {

		return 0, err
	}

	return Response.Count, nil
}

/**
 * @description: 向某个client发送消息
 * @param {string} clientid
 * @param {string} message
 * @return {*}
 */
func (c *GatewayWebSocketClient) SendMessageToClient(clientid string, message string) (bool, error) {

	Request := &SendMessageToClientPB.SendMessageToClientRequest{}
	Request.Clientid = clientid
	Request.Message = message

	Response, err := SendMessageToClientPB.NewSendMessageToClientClient(c.conn).SendMessageToClient(context.Background(), Request)

	if err != nil {

		return false, err
	}

	return Response.Result, nil
}

/**
 * @description: 获取client绑定的Uid
 * @param {string} clientid
 * @return {*}
 */
func (c *GatewayWebSocketClient) GetClientByUid(clientid string) (string, error) {

	Request := &GetClientByUidPB.GetClientByUidRequest{}
	Request.Clientid = clientid

	Response, err := GetClientByUidPB.NewGetClientByUidClient(c.conn).GetClientByUid(context.Background(), Request)

	if err != nil {

		return "", nil
	}

	return Response.Uid, nil
}

/**
 * @description:获取所有在线客户端
 * @return {*}
 */
func (c *GatewayWebSocketClient) GetAllOnlineClient() ([]string, error) {

	Request := &GetAllOnlineClientPB.GetAllOnlineClientRequest{}

	Response, err := GetAllOnlineClientPB.NewGetAllOnlineClientClient(c.conn).GetAllOnlineClient(context.Background(), Request)

	if err != nil {

		return []string{}, err
	}

	return Response.Client, nil
}

/**
 * @description: 统计在线client数量
 * @return {*}
 */
func (c *GatewayWebSocketClient) CountOnlineClient() (int32, error) {

	Request := &CountOnlineClientPB.CountOnlineClientRequest{}

	Response, err := CountOnlineClientPB.NewCountOnlineClientClient(c.conn).CountOnlineClient(context.Background(), Request)

	if err != nil {

		return 0, err
	}

	return Response.Count, nil
}

/**
 * @description: 关闭某个client的连接
 * @param {string} clientid
 * @return {*}
 */
func (c *GatewayWebSocketClient) ClonseClient(clientid string) (bool, error) {

	Request := &ClonseClientPB.ClonseClientRequest{}
	Request.Clientid = clientid

	Response, err := ClonseClientPB.NewClonseClientClient(c.conn).ClonseClient(context.Background(), Request)

	if err != nil {

		return false, err
	}

	return Response.Result, nil
}

/**
 * @description: 判断client是否在线
 * @param {string} clientid
 * @return {*}
 */
func (c *GatewayWebSocketClient) ClientIsOnline(clientid string) (bool, error) {

	Request := &ClientIsOnlinePB.ClientIsOnlineRequest{}
	Request.Clientid = clientid

	Response, err := ClientIsOnlinePB.NewClientIsOnlineClient(c.conn).ClientIsOnline(context.Background(), Request)

	if err != nil {

		return false, err
	}

	return Response.Isonline, nil
}

/**
 * @description: client绑定Uid
 * @param {string} clientid
 * @param {string} uid
 * @return {*}
 */
func (c *GatewayWebSocketClient) ClientBindUid(clientid string, uid string) (bool, error) {

	Request := &ClientBindUidPB.ClientBindUidRequest{}
	Request.Clientid = clientid
	Request.Uid = uid

	Response, err := ClientBindUidPB.NewClientBindUidClient(c.conn).ClientBindUid(context.Background(), Request)

	if err != nil {

		return false, err
	}

	return Response.Result, nil
}

/**
 * @description: 向所有连接广播一条相同的消息
 * @param {string} message
 * @return {*}
 */
func (c *GatewayWebSocketClient) BroadcastMessage(message string) (bool, error) {

	Request := &BroadcastMessagePB.BroadcastMessageRequest{}
	Request.Message = message

	Response, err := BroadcastMessagePB.NewBroadcastMessageClient(c.conn).BroadcastMessage(context.Background(), Request)

	if err != nil {

		return false, err
	}

	return Response.Result, nil
}

/**
 * @description: 初始化gRPC连接
 * @param {string} address
 * @return {*}
 */
func NewGRPCClient(address string) (*GatewayWebSocketClient, error) {

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("failed to connect to gRPC server: %v", err)
	}

	client := &GatewayWebSocketClient{
		conn: conn,
	}

	runtime.SetFinalizer(client, finalizeGRPCClient)

	return client, nil
}

/**
 * @description: 调用结束后自动关闭gRPC连接
 * @param {*GatewayWebSocketClient} c
 * @return {*}
 */
func finalizeGRPCClient(c *GatewayWebSocketClient) {
	c.conn.Close()
}
