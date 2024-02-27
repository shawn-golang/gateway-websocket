/*
 * @Author: psq
 * @Date: 2023-05-09 10:00:18
 * @LastEditors: psq
 * @LastEditTime: 2024-02-27 11:42:23
 */

package command

import (
	"fmt"
	"gateway-websocket/config"
	"io/ioutil"
	"net"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"

	gRPCClient "gateway-websocket/services/client"
)

var (
	pidFile = "gateway-websocket.pid"
	version = "1.2.0.1"
)

func getGatewayPID() string {

	data, err := ioutil.ReadFile(pidFile)
	if err != nil {

		return ""
	}

	return string(data)
}

func GatewayStatus() {

	if _, err := os.Stat(getGatewayPID()); err == nil {

		fmt.Println("gateway-websocket service not started.")
		return
	}

	client, err := gRPCClient.NewGRPCClient(fmt.Sprintf("localhost:%d", config.GatewayConfig["GRPCServicePort"]))

	if err != nil {

		fmt.Println("gateway-websocket gRPC communication failure.")
		return
	}

	countOnlineUser, countOnlineClient, countOnlineGroup := int32(0), int32(0), int32(0)

	res, err := client.CountOnlineUid()

	if err == nil {

		countOnlineUser = res
	}

	res, err = client.CountOnlineClient()

	if err == nil {

		countOnlineClient = res
	}

	res, err = client.CountGroup()

	if err == nil {

		countOnlineGroup = res
	}

	usage := "----------------------------------------------Gateway-Websocket Status----------------------------------------------------"

	// 获取进程信息
	cmd, err := exec.Command("ps", "-p", getGatewayPID(), "-o", "%cpu,%mem,etime=,lstart=").Output()
	if err != nil {

		fmt.Println("gateway-websocket service not started.")
		return
	}

	for _, line := range strings.Split(string(cmd), "\n")[1:] {

		fields := strings.Fields(line)

		if len(fields) > 0 {

			usage += fmt.Sprintf("\nGateway version: %s\tCPU Use: %s\tMemory Use: %s\tTime: %s\tStarted at: %s", version, fields[0], fields[1], fields[2], strings.Join(fields[3:], " "))
		}
	}

	usage += fmt.Sprintf("\nGateway port: %d \t\tgRPC port: %d\n", config.GatewayConfig["GatewayServicePort"], config.GatewayConfig["GRPCServicePort"])

	usage += "----------------------------------------------Gateway-Websocket Connect----------------------------------------------------\n"

	usage += fmt.Sprintf("Client: %d\tUsers: %d\tGroup:  %d\n", countOnlineClient, countOnlineUser, countOnlineGroup)

	// usage += "----------------------------------------------Gateway-Websocket Cluster----------------------------------------------------"

	fmt.Println(usage)
}

func StartGateway() {

	fmt.Println("Start gateway-websocket server...")

	// 检查gRPC端口占用
	if _, err := net.Listen("tcp", ":"+fmt.Sprintf("%v", config.GatewayConfig["GRPCServicePort"])); err != nil {

		fmt.Printf("gRPC service Port %s is already in use\n", fmt.Sprintf("%v", config.GatewayConfig["GRPCServicePort"]))
		return
	}

	// 检查WebSocket端口占用
	if _, err := net.Listen("tcp", ":"+fmt.Sprintf("%v", config.GatewayConfig["GatewayServicePort"])); err != nil {

		fmt.Printf("WebSocket ervice Port %s is already in use\n", fmt.Sprintf("%v", config.GatewayConfig["GatewayServicePort"]))
		return
	}

	// 以守护进程方式启动程序
	cmd := exec.Command(os.Args[0], "daemon")
	cmd.SysProcAttr = &syscall.SysProcAttr{Setsid: true}

	err := cmd.Start()
	if err != nil {

		fmt.Println(err)
		return
	}

	err = ioutil.WriteFile(pidFile, []byte(strconv.Itoa(cmd.Process.Pid)), 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("successful, gateway-websocket listening on port", config.GatewayConfig["GatewayServicePort"])
}

func StopGateway() {

	pid, err := strconv.Atoi(getGatewayPID())
	if err != nil {
		fmt.Println(err)
		return
	}

	process, err := os.FindProcess(pid)
	if err != nil {
		fmt.Println(err)
		return
	}

	if err = process.Kill(); err != nil {

		fmt.Println(err)
		return
	}
}
