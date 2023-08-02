/*
 * @Author: psq
 * @Date: 2023-04-24 15:14:11
 * @LastEditors: psq
 * @LastEditTime: 2023-08-02 18:04:41
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

	go gRPC.RegService()

	engine := gin.Default()

	engine.GET("/", WebSocket.WsServer)

	engine.Run(fmt.Sprintf(":%d", config.GatewayConfig["GatewayServicePort"]))
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
	 `
	fmt.Println(usage)
}
