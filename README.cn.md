
# gateway-websocket
[![Go](https://img.shields.io/badge/Go->=1.17-green)](https://go.dev)
[![Release](https://img.shields.io/github/v/release/jefferyjob/go-easy-utils.svg)](https://github.com/jefferyjob/go-easy-utils/releases)
[![Action](https://github.com/jefferyjob/go-easy-utils/workflows/Go/badge.svg?branch=main)](https://github.com/jefferyjob/go-easy-utils/actions)
[![Report](https://goreportcard.com/badge/github.com/jefferyjob/go-easy-utils)](https://goreportcard.com/report/github.com/jefferyjob/go-easy-utils)
[![Coverage](https://codecov.io/gh/jefferyjob/go-easy-utils/branch/main/graph/badge.svg)](https://codecov.io/gh/jefferyjob/go-easy-utils)
[![Doc](https://img.shields.io/badge/go.dev-reference-brightgreen?logo=go&logoColor=white&style=flat)](https://pkg.go.dev/github.com/jefferyjob/go-easy-utils/v2)
[![License](https://img.shields.io/github/license/jefferyjob/go-easy-utils)](https://github.com/jefferyjob/go-easy-utils/blob/main/LICENSE)

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
        gateway-websocket在设计之初的宗旨就是不侵入业务，这里的 ##user## 只是开发者对一个或多个client的通称，需要注意的是，一个client只能绑定一个uid，但一个uid可以绑定多个client
    
    # group
        和user类似，group是用来将不同的client加入到某群组中进行消息通讯


