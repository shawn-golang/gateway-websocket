/*
 * @Author: psq
 * @Date: 2023-05-09 10:00:18
 * @LastEditors: psq
 * @LastEditTime: 2023-12-12 15:56:22
 */

package command

import (
	"context"
	"fmt"
	"gateway-websocket/config"
	CountGroupPB "gateway-websocket/services/grpc/proto/countgroup"
	CountOnlineClientPB "gateway-websocket/services/grpc/proto/countonlineclient"
	CountOnlineUidPB "gateway-websocket/services/grpc/proto/countonlineuid"
	"io/ioutil"
	"net"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"

	"google.golang.org/grpc"
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

	pid := getGatewayPID()

	if _, err := os.Stat(pid); err == nil {

		fmt.Println("gateway-websocket service not started.")
		return
	}

	conn, err := grpc.Dial(fmt.Sprintf("localhost:%d", config.GatewayConfig["GRPCServicePort"]), grpc.WithInsecure())

	if err != nil {

		fmt.Println("gateway-websocket gRPC communication failure.")
		return
	}

	defer conn.Close()

	countOnlineUser, countOnlineClient, countOnlineGroup := 0, 0, 0

	usersRequest := &CountOnlineUidPB.CountOnlineUidRequest{}

	usersResponse, usersErr := CountOnlineUidPB.NewCountOnlineUidClient(conn).CountOnlineUid(context.Background(), usersRequest)

	if usersErr == nil {

		countOnlineUser = int(usersResponse.Count)
	}

	clientRequest := &CountOnlineClientPB.CountOnlineClientRequest{}

	clientResponse, clientErr := CountOnlineClientPB.NewCountOnlineClientClient(conn).CountOnlineClient(context.Background(), clientRequest)

	if clientErr == nil {

		countOnlineClient = int(clientResponse.Count)
	}

	groupRequest := &CountGroupPB.CountGroupRequest{}

	groupResponse, groupErr := CountGroupPB.NewCountGroupClient(conn).CountGroup(context.Background(), groupRequest)

	if groupErr == nil {

		countOnlineGroup = int(groupResponse.Count)
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
