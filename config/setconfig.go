/*
 * @Author: psq
 * @Date: 2023-04-27 18:39:35
 * @LastEditors: psq
 * @LastEditTime: 2023-05-08 20:17:12
 */
package config

import (
	"encoding/json"
	"net"
	"os"
	"strconv"

	"gopkg.in/ini.v1"
)

var GatewayConfig = make(map[string]interface{})

func init() {

	GatewayConfig["GatewayServicePort"] = 20818
	GatewayConfig["ReadBufferSize"] = 1024
	GatewayConfig["WriteBufferSize"] = 1024
	GatewayConfig["MessageCompression"] = true
	GatewayConfig["HeartbeatTimeout"] = 180
	GatewayConfig["WebsocketHandshakeTimeout"] = 5
	GatewayConfig["WebsocketRoute"] = "wss"
	GatewayConfig["GRPCServicePort"] = 20819
	GatewayConfig["ClientIP"] = []string{}

	configPath := "./config/config.ini"

	checkPath := func(path string) bool {

		_, err := os.Stat(path)

		return err == nil || os.IsExist(err)
	}

	if !checkPath(configPath) {

		return
	}

	cnf, err := ini.Load("./config/config.ini")

	if err != nil {

		return
	}

	// 设置gRPC服务端运行端口
	if cnf.Section("gateway").Key("gRPCServerPort").String() != "" {

		gRPCServicePort, err := strconv.Atoi(cnf.Section("gateway").Key("gRPCServerPort").String())

		if err == nil {

			GatewayConfig["GRPCServicePort"] = gRPCServicePort
		}
	}

	// 设置ws服务端运行端口
	if cnf.Section("gateway").Key("gatewayServicePort").String() != "" {

		gatewayServicePort, err := strconv.Atoi(cnf.Section("gateway").Key("gatewayServicePort").String())

		if err == nil {

			GatewayConfig["GatewayServicePort"] = gatewayServicePort
		}
	}

	// 消息接收缓冲区大小
	if cnf.Section("gateway").Key("readBufferSize").String() != "" {

		readBufferSize, err := strconv.Atoi(cnf.Section("gateway").Key("readBufferSize").String())

		if err == nil {

			GatewayConfig["ReadBufferSize"] = readBufferSize
		}
	}

	// 消息发送缓冲区大小
	if cnf.Section("gateway").Key("writeBufferSize").String() != "" {

		writeBufferSize, err := strconv.Atoi(cnf.Section("gateway").Key("writeBufferSize").String())

		if err == nil {

			GatewayConfig["WriteBufferSize"] = writeBufferSize
		}
	}

	// 是否压缩消息包
	if cnf.Section("gateway").Key("writeBufferSize").String() != "" {

		messageCompression := cnf.Section("gateway").Key("messageCompression").String()

		if messageCompression == "true" {

			GatewayConfig["MessageCompression"] = true
		}

		if messageCompression == "false" {

			GatewayConfig["MessageCompression"] = false
		}
	}

	// 心跳超时时间
	if cnf.Section("gateway").Key("heartbeatTimeout").String() != "" {

		heartbeatTimeout, err := strconv.Atoi(cnf.Section("gateway").Key("heartbeatTimeout").String())

		if err == nil {

			GatewayConfig["HeartbeatTimeout"] = heartbeatTimeout
		}
	}

	// 设置ws握手超时时间
	if cnf.Section("gateway").Key("websocketHandshakeTimeout").String() != "" {

		websocketHandshakeTimeout, err := strconv.Atoi(cnf.Section("gateway").Key("websocketHandshakeTimeout").String())

		if err == nil {

			GatewayConfig["WebsocketHandshakeTimeout"] = websocketHandshakeTimeout
		}
	}

	// 设置ws连接路由
	if cnf.Section("gateway").Key("websocketRoute").String() != "" {

		websocketRoute, err := strconv.Atoi(cnf.Section("gateway").Key("websocketRoute").String())

		if err == nil {

			GatewayConfig["WebsocketRoute"] = websocketRoute
		}
	}

	if cnf.Section("client").Key("ip").String() != "" {

		var ips []string

		if err := json.Unmarshal([]byte(cnf.Section("client").Key("ip").String()), &ips); err != nil {

			return
		}

		var validIPs []string

		for _, ip := range ips {

			if net.ParseIP(ip) != nil {

				validIPs = append(validIPs, ip)
			}
		}

		GatewayConfig["ClientIP"] = validIPs
	}
}
