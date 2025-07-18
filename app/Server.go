package main

import (
	"fmt"

	"github.com/ibuprofen/Tin/tinface"
	"github.com/ibuprofen/Tin/tnet"
)

// ping test 自定义路由
type PingRouter struct {
	tnet.BaseRouter // 一定要先基础BaseRouter
}

// Test PreHandle
func (pr *PingRouter) PreHandle(request tinface.IRequest) {
	fmt.Println("Call Router PreHandle")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("before ping ....\n"))
	if err != nil {
		fmt.Println("call back ping ping ping error")
	}
}

// Test Handle
func (pr *PingRouter) Handle(request tinface.IRequest) {
	fmt.Println("Call PingRouter Handle")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("ping...ping...ping\n"))
	if err != nil {
		fmt.Println("call back ping ping ping error")
	}
}

// Test PostHandle
func (pr *PingRouter) PostHandle(request tinface.IRequest) {
	fmt.Println("Call Router PostHandle")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("After ping .....\n"))
	if err != nil {
		fmt.Println("call back ping ping ping error")
	}
}

// Server 模块的测试函数
func main() {
	// 1 创建一个server 句柄 s
	s := tnet.NewServer()

	s.AddRouter(&PingRouter{})

	// 2 开启服务
	s.Serve()
}
