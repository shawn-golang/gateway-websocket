<!--
 * @Author: psq
 * @Date: 2023-12-11 14:10:31
 * @LastEditors: psq
 * @LastEditTime: 2023-12-13 17:18:34
-->
# gateway-websocket
[![Go](https://img.shields.io/badge/Go->=1.17-green)](https://go.dev)
<!-- [![Release](https://img.shields.io/github/v/release/jefferyjob/go-easy-utils.svg)](https://github.com/jefferyjob/go-easy-utils/releases)
[![Action](https://github.com/jefferyjob/go-easy-utils/workflows/Go/badge.svg?branch=main)](https://github.com/jefferyjob/go-easy-utils/actions)
[![Report](https://goreportcard.com/badge/github.com/jefferyjob/go-easy-utils)](https://goreportcard.com/report/github.com/jefferyjob/go-easy-utils)
[![Coverage](https://codecov.io/gh/jefferyjob/go-easy-utils/branch/main/graph/badge.svg)](https://codecov.io/gh/jefferyjob/go-easy-utils)
[![Doc](https://img.shields.io/badge/go.dev-reference-brightgreen?logo=go&logoColor=white&style=flat)](https://pkg.go.dev/github.com/jefferyjob/go-easy-utils/v2)
[![License](https://img.shields.io/github/license/jefferyjob/go-easy-utils)](https://github.com/jefferyjob/go-easy-utils/blob/main/LICENSE) -->

[English](README.md) | [简体中文](README.cn.md)

## gateway-websocket介绍
gateway-websocket是一个基于golang开发的websocket服务端，旨在为开发者提供一个高效率、高稳定性、对业务系统零侵入的开源websocket服务器

## 快速开始

```bash

go run ./main.go start  // 启动gateway-websocket服务器
go run ./main.go stop   // 停止gateway-websocket服务器
go run ./main.go status // 查看gateway-websocket服务器运行状态

go run ./main.go reload // 热更新gateway-websocket配置【功能正在开发】
```

## 使用说明

gateway-websocket中有三个重要的概念 `client-客户端`，`user-用户`，`group-群组`

    # client
        gateway-websocket的概念中任意一个设备只要通过ws/wss连接到服务端时就会被认定为一个有效的 #client#
    
    # user
        gateway-websocket在设计之初的宗旨就是不侵入业务，这里的 ##user## 只是开发者对一个或多个client的统称，需要注意的是，一个client只能绑定一个uid，但一个uid可以绑定多个client
    
    # group
        和user类似，group是用来将不同的client加入到某群组中进行消息通讯


## 调用gateway-websocket提供的接口

gateway-websocket在设计之初就考虑到了其他开发语言调用接口的问题，所以gateway-websocket为大家提供了`gRPC`接口

| 接口名称 | 用途                                                                         | 归属分类             |
|--------------| ----------------------------------------------------------------------------------------- |----------------------|
| BroadcastMessage      | 向所有client广播消息                                           | client    |
| SendMessageToClient     | 向某个client发送消息                                                                  | client   |
| ClonseClient   | 关闭某个client的连接                                                             | client |
| GetAllOnlineClient    | 获取所有在线客户端                                                 | client  |
| ClientIsOnline    | 判断client是否在线                                                            | client  |
| CountOnlineClient      | 统计在线client数量                                                                 | client   |
| UnGroup     | 解散某个群组                                        | group   |
| SendMessageToGroup      | 向某个群组发送消息                                                                  | group    |
| CountOnlineGroup     | 统计群组内在线的client数量              | group   |
| CountGroup     | 统计群组数量                           | group   |
| GetGroupOnlineClient    | 获取群组内所有在线的client                     | group  |
| LeaveGroup      | 将client移出群组                                                              | group   |
| JoinGroup    | 将client加入群组 | group  |
| GetUidByClient    | 获取Uid下所有client | user  |
| CountOnlineUid    | 统计在线uid人数 | user  |
| GetAllOnlineUid    | 获取所有在线uid | user  |
| SendMessageToUid    | 向某个uid下绑定的所有client发送消息 | user  |
| UidIsOnline    | 判断某个uid是否存在 | user  |
| UnBindUid    | client解除绑定Uid | user  |
| ClientBindUid    | client绑定Uid | user  |
| GetClientByUid    | 获取client绑定的Uid | user  |

# gateway-client 调用示例

[各语言调用示例请参阅（逐步完善中）](https://github.com/shawn-golang/gateway-websocket/releases)
