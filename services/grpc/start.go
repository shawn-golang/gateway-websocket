/*
 * @Author: psq
 * @Date: 2023-04-28 20:30:51
 * @LastEditors: psq
 * @LastEditTime: 2023-08-02 18:05:31
 */
package services

import (
	"context"
	"fmt"
	"gateway-websocket/config"
	"net"

	Controllers "gateway-websocket/services/grpc/controllers"
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

	"google.golang.org/grpc"
	GRPC "google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

/**
 * @description: gRPC拦截器
 * @param {context.Context} ctx
 * @param {interface{}} req
 * @param {*grpc.UnaryServerInfo} info
 * @param {grpc.UnaryHandler} handler
 * @return {*}
 */
func identityCheck(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {

	p, _ := peer.FromContext(ctx)

	// 拆分地址信息，获取 IP 地址部分
	host, _, _ := net.SplitHostPort(p.Addr.String())

	pass := false

	ipaddr := config.GatewayConfig["ServerIP"].([]string)

	if len(ipaddr) >= 1 {

		for _, v := range ipaddr {

			if host == v {

				pass = true
				break
			}
		}
	} else {
		pass = true
	}

	if !pass {

		return nil, status.Errorf(codes.InvalidArgument, "invalid request")
	}

	return handler(ctx, req)
}

func RegService() {

	listenPort, _ := net.Listen("tcp", fmt.Sprintf(":%d", config.GatewayConfig["GRPCServicePort"]))

	gRPC := GRPC.NewServer(GRPC.UnaryInterceptor(identityCheck))

	reflection.Register(gRPC)

	// client相关接口
	CountOnlineClientPB.RegisterCountOnlineClientServer(gRPC, &Controllers.ClientControllers{})
	ClientIsOnlinePB.RegisterClientIsOnlineServer(gRPC, &Controllers.ClientControllers{})
	GetAllOnlineClientPB.RegisterGetAllOnlineClientServer(gRPC, &Controllers.ClientControllers{})
	ClonseClientPB.RegisterClonseClientServer(gRPC, &Controllers.ClientControllers{})
	SendMessageToClientPB.RegisterSendMessageToClientServer(gRPC, &Controllers.ClientControllers{})
	BroadcastMessagePB.RegisterBroadcastMessageServer(gRPC, &Controllers.ClientControllers{})

	// users相关接口
	GetClientByUidPB.RegisterGetClientByUidServer(gRPC, &Controllers.UserControllers{})
	ClientBindUidPB.RegisterClientBindUidServer(gRPC, &Controllers.UserControllers{})
	CountOnlineUidPB.RegisterCountOnlineUidServer(gRPC, &Controllers.UserControllers{})
	GetAllOnlineUidPB.RegisterGetAllOnlineUidServer(gRPC, &Controllers.UserControllers{})
	UnBindUidPB.RegisterUnBindUidServer(gRPC, &Controllers.UserControllers{})
	UidIsOnlinePB.RegisterUidIsOnlineServer(gRPC, &Controllers.UserControllers{})
	SendMessageToUidPB.RegisterSendMessageToUidServer(gRPC, &Controllers.UserControllers{})
	GetUidByClientPB.RegisterGetUidByClientServer(gRPC, &Controllers.UserControllers{})

	// group相关接口
	CountGroupPB.RegisterCountGroupServer(gRPC, &Controllers.GroupControllers{})
	CountOnlineGroupPB.RegisterCountOnlineGroupServer(gRPC, &Controllers.GroupControllers{})
	GetGroupOnlineClientPB.RegisterGetGroupOnlineClientServer(gRPC, &Controllers.GroupControllers{})
	JoinGroupPB.RegisterJoinGroupServer(gRPC, &Controllers.GroupControllers{})
	LeaveGroupPB.RegisterLeaveGroupServer(gRPC, &Controllers.GroupControllers{})
	SendMessageToGroupPB.RegisterSendMessageToGroupServer(gRPC, &Controllers.GroupControllers{})
	UnGroupPB.RegisterUnGroupServer(gRPC, &Controllers.GroupControllers{})

	// 启动服务
	gRPC.Serve(listenPort)
}
