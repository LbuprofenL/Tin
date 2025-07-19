package main

import (
	"fmt"

	"github.com/ibuprofen/Tin/tinface"
	"github.com/ibuprofen/Tin/tnet"
)

// ping test 自定义路由
type PingRouter struct {
	tnet.BaseRouter
}

// Ping Handle
func (this *PingRouter) Handle(request tinface.IRequest) {
	fmt.Println("Call PingRouter Handle")
	// 先读取客户端的数据，再回写ping...ping...ping
	fmt.Println("recv from client : msgId=", request.GetMsgID(), ", data=", string(request.GetData()))

	err := request.GetConnection().SendMsg(0, []byte("ping...ping...ping"))
	if err != nil {
		fmt.Println(err)
	}
}

// HelloZinxRouter Handle
type HelloZinxRouter struct {
	tnet.BaseRouter
}

func (this *HelloZinxRouter) Handle(request tinface.IRequest) {
	fmt.Println("Call HelloZinxRouter Handle")
	// 先读取客户端的数据，再回写ping...ping...ping
	fmt.Println("recv from client : msgId=", request.GetMsgID(), ", data=", string(request.GetData()))

	err := request.GetConnection().SendMsg(1, []byte("Hello Tin Router V0.6"))
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	// 创建一个server句柄
	s := tnet.NewServer()

	// 配置路由
	s.AddRouter(0, &PingRouter{})
	s.AddRouter(1, &HelloZinxRouter{})

	// 开启服务
	s.Serve()
}
