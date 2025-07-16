package main

import (
	"github.com/ibuprofen/Tin/tin/tnet"
)

// Server 模块的测试函数
func main() {
	// 1 创建一个server 句柄 s
	s := tnet.NewServer("[tin V0.1]")

	// 2 开启服务
	s.Serve()
}
