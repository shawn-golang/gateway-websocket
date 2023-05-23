/*
 * @Author: psq
 * @Date: 2023-04-24 15:14:11
 * @LastEditors: psq
 * @LastEditTime: 2023-05-12 15:50:50
 */
package main

import (
	"fmt"
	"gateway-websocket/command"
	"gateway-websocket/config"
	gRPC "gateway-websocket/services/grpc"
	WebSocket "gateway-websocket/services/websocket"
	"os"

	"github.com/gin-gonic/gin"
)

func gatewayWebsocketService() {

	gin.SetMode(gin.ReleaseMode)

	go gRPC.GrpcServiceStart()

	engine := gin.Default()

	engine.GET("/"+fmt.Sprintf("%v", config.GatewayConfig["WebsocketRoute"]), WebSocket.WsServer)

	engine.Run(":" + fmt.Sprintf("%v", config.GatewayConfig["GatewayServicePort"]))
}

func main() {

	args := os.Args[1:]

	if len(args) > 0 {

		switch args[0] {

		case "start":

			command.StartService()
			return
		case "stop":

			command.StopService()
			return
		case "status":

			command.DumpServieStatus()
			return
		case "reload":
		case "daemon":

			gatewayWebsocketService()
			return

		default:

			goto TIPS
		}

	} else {

		goto TIPS
	}

TIPS:

	usage := `
gateway-websocket is an efficient WebSocket server.

Usage:

    gateway-websocket <command>

The commands are:

    start       starts the gateway-websocket process
    stop        stops the gateway-websocket process
    status      displays the runtime status and client connection information of gateway-websocket
    reload      hot-reloads the configuration file
    `
	fmt.Println(usage)
}
