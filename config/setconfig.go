/*
 * @Author: psq
 * @Date: 2023-04-27 18:39:35
 * @LastEditors: psq
 * @LastEditTime: 2023-12-12 16:01:41
 */
package config

import (
	"encoding/json"
	"net"
	"os"
	"strconv"

	"github.com/gorilla/websocket"
	"gopkg.in/ini.v1"
)

var GatewayConfig = make(map[string]interface{})

func init() {

	GatewayConfig["GatewayServicePort"] = 20818
	GatewayConfig["WriteBufferSize"] = 1024
	GatewayConfig["HeartbeatTimeout"] = 180
	GatewayConfig["GRPCServicePort"] = 20819
	GatewayConfig["ClientIP"] = []string{}
	GatewayConfig["MessageFormat"] = websocket.TextMessage

	configPath := "./config/config.ini"

	checkPath := func(path string) bool {

		_, err := os.Stat(path)

		return err == nil || os.IsExist(err)
	}

	if !checkPath(configPath) {

		return
	}

	cnf, err := ini.Load(configPath)

	if err != nil {

		return
	}

	// 设置gRPC服务端运行端口
	if cnf.Section("gateway").Key("gRPCServicePort").String() != "" {

		gRPCServicePort, err := strconv.Atoi(cnf.Section("gateway").Key("gRPCServicePort").String())

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

	// 消息发送缓冲区大小
	if cnf.Section("gateway").Key("writeBufferSize").String() != "" {

		writeBufferSize, err := strconv.Atoi(cnf.Section("gateway").Key("writeBufferSize").String())

		if err == nil {

			GatewayConfig["WriteBufferSize"] = writeBufferSize
		}
	}

	// 设置websocket消息体格式
	if cnf.Section("gateway").Key("messageFormat").String() != "" {

		if cnf.Section("gateway").Key("messageFormat").String() == "binary" {

			GatewayConfig["MessageFormat"] = websocket.BinaryMessage

		} else {

			GatewayConfig["MessageFormat"] = websocket.TextMessage
		}
	}

	// 心跳超时时间
	if cnf.Section("gateway").Key("heartbeatTimeout").String() != "" {

		heartbeatTimeout, err := strconv.Atoi(cnf.Section("gateway").Key("heartbeatTimeout").String())

		if err == nil {

			GatewayConfig["HeartbeatTimeout"] = heartbeatTimeout
		}
	}

	if cnf.Section("server").Key("ip").String() != "" {

		var ips []string

		if err := json.Unmarshal([]byte(cnf.Section("server").Key("ip").String()), &ips); err != nil {

			return
		}

		var validIPs []string

		for _, ip := range ips {

			if net.ParseIP(ip) != nil {

				validIPs = append(validIPs, ip)
			}
		}

		GatewayConfig["ServerIP"] = validIPs
	}
}
