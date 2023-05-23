/*
 * @Author: psq
 * @Date: 2023-05-09 10:00:18
 * @LastEditors: psq
 * @LastEditTime: 2023-05-22 18:01:48
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
	pidFile  = "gateway-websocket.pid"
	version  = "1.0.0.1"
	gRPCPort = config.GatewayConfig["GRPCServicePort"]
)

func countOnlineCient() int {

	conn, err := grpc.Dial(fmt.Sprintf("localhost:%d", gRPCPort), grpc.WithInsecure())
	if err != nil {

		return 0
	}

	defer conn.Close()

	c := CountOnlineClientPB.NewCountOnlineClientClient(conn)

	req := &CountOnlineClientPB.CountOnlineClientRequest{}

	res, err := c.CountOnlineClient(context.Background(), req)

	if err != nil {
		fmt.Println(err)
		return 0
	}

	return int(res.Count)
}

func countOnlineGroup() int {

	conn, err := grpc.Dial(fmt.Sprintf("localhost:%d", gRPCPort), grpc.WithInsecure())

	if err != nil {

		return 0
	}

	defer conn.Close()

	c := CountGroupPB.NewCountGroupClient(conn)

	req := &CountGroupPB.CountGroupRequest{}

	res, err := c.CountGroup(context.Background(), req)

	if err != nil {

		return 0
	}

	return int(res.Count)

}

func countOnlineUser() int {

	conn, err := grpc.Dial(fmt.Sprintf("localhost:%d", gRPCPort), grpc.WithInsecure())
	if err != nil {

		return 0
	}

	defer conn.Close()

	c := CountOnlineUidPB.NewCountOnlineUidClient(conn)

	req := &CountOnlineUidPB.CountOnlineUidRequest{}

	res, err := c.CountOnlineUid(context.Background(), req)

	if err != nil {

		return 0
	}

	return int(res.Count)

}

func getGatewayPID() string {

	data, err := ioutil.ReadFile(pidFile)
	if err != nil {

		return ""
	}

	return string(data)
}

func DumpServieStatus() {

	pid := getGatewayPID()
	if _, err := os.Stat(pid); err == nil {
		fmt.Println("gateway-websocket service not started.")
		return
	}

	usage := "----------------------------------------------Gatewat-Websocket Status----------------------------------------------------"

	// 获取进程信息
	cmd, err := exec.Command("ps", "-p", getGatewayPID(), "-o", "%cpu,%mem,etime=,lstart=").Output()
	if err != nil {

		fmt.Println("gateway-websocket service not started.")
		return
	}

	for _, line := range strings.Split(string(cmd), "\n")[1:] {

		fields := strings.Fields(line)

		if len(fields) > 0 {

			usage += fmt.Sprintf("\nGateway version: %s\tCPU: %s\tMemory: %s\tTime: %s\tStarted at: %s", version, fields[0], fields[1], fields[2], strings.Join(fields[3:], " "))
		}
	}

	usage += fmt.Sprintf("\nGateway port: %d \t\tgRPC port: %d\n", config.GatewayConfig["GatewayServicePort"], config.GatewayConfig["GRPCServicePort"])

	usage += "----------------------------------------------Gatewat-Websocket Connect----------------------------------------------------\n"

	usage += fmt.Sprintf("Client: %d\tRegister Users: %d\tGroup:  %d\n", countOnlineCient(), countOnlineUser(), countOnlineGroup())

	usage += "----------------------------------------------Gatewat-Websocket Cluster----------------------------------------------------"

	fmt.Println(usage)
}

func StartService() {

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

	cmd := exec.Command(os.Args[0], "daemon")
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Setsid: true,
	}

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

	fmt.Println("Start gateway-websocket server...")
	fmt.Println("successful，gateway-websocket listening on port", config.GatewayConfig["GatewayServicePort"])
}

func StopService() {

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
