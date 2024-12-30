/*
 * @Author: psq
 * @Date: 2023-04-24 15:14:11
 * @LastEditors: psq
 * @LastEditTime: 2024-12-30 13:56:15
 */
package main

import (
	"fmt"
	"gateway-websocket/command"
	"gateway-websocket/config"
	gRPC "gateway-websocket/services/grpc"
	webSocket "gateway-websocket/services/websocket"
	"net/http"
	"os"
)

/**
 * @description: 启动GatewayWebsocket服务进程
 * @return {*}
 */
func gatewayWebsocketService() {

	go gRPC.RegService()

	// 使用net/http包运行socket进程
	http.HandleFunc("/", webSocket.WsServer)

	err := http.ListenAndServe(fmt.Sprintf(":%d", config.GatewayConfig["GatewayServicePort"]), nil)

	if err != nil {

		fmt.Println("Error starting server:", err)
	}
	// 使用gin包运行socket进程
	// gin.SetMode(gin.ReleaseMode)

	// engine := gin.Default()

	// engine.GET("/", webSocket.WsServer)

	// engine.Run(fmt.Sprintf(":%d", config.GatewayConfig["GatewayServicePort"]))

	// 上述web包按照你的开发习惯选择其中一种即可，当然你也可以用别的web包
}

func main() {

	args := os.Args[1:]

	if len(args) <= 0 || len(args) > 1 {

		goto TIPS
	}

	switch args[0] {

	case "start":
		command.StartGateway()
		return
	case "stop":

		command.StopGateway()
		return
	case "status":

		command.GatewayStatus()
		return
	case "daemon":

		gatewayWebsocketService()
		return

	default:

		goto TIPS
	}

TIPS:

	fmt.Println(`
	gateway-websocket is an efficient WebSocket server.

	Usage:

		gateway-websocket <command>

	The commands are:

		start       starts the gateway-websocket process
		stop        stops the gateway-websocket process
		status      displays the runtime status and client connection information of gateway-websocket
		`)
}
