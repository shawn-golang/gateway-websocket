/*
 * @Author: psq
 * @Date: 2023-04-24 15:14:11
 * @LastEditors: psq
 * @LastEditTime: 2024-02-27 13:58:17
 */
package main

import (
	"fmt"
	"gateway-websocket/command"
	"gateway-websocket/config"
	gRPC "gateway-websocket/services/grpc"
	webSocket "gateway-websocket/services/websocket"
	"os"

	"github.com/gin-gonic/gin"
)

func gatewayWebsocketService() {

	go gRPC.RegService()

	// http.HandleFunc("/", webSocket.WsServer)

	// err := http.ListenAndServe(fmt.Sprintf(":%d", config.GatewayConfig["GatewayServicePort"]), nil)

	// if err != nil {

	// 	fmt.Println("Error starting server:", err)
	// }

	gin.SetMode(gin.ReleaseMode)

	engine := gin.Default()

	engine.GET("/", webSocket.WsServer)

	engine.Run(fmt.Sprintf(":%d", config.GatewayConfig["GatewayServicePort"]))
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
