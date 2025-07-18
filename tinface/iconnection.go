package tinface

import "net"

// IConnection 定义一个服务器接口
type IConnection interface {
	// 启动连接，让当前连接开始工作
	Start()
	// 停止连接，结束当前连接状态M
	Stop()
	// 从当前连接获取原始的socket TCPConn
	GetTCPConnection() *net.TCPConn
	// 获取当前连接ID
	GetConnID() uint32
	// 获取远程客户端地址信息
	RemoteAddr() net.Addr
	// 发送数据，将数据发送给远程的客户端
	SendMsg(msgID uint32, data []byte) error
}

// HandFunc 定义一个统一处理链接业务的接口
type HandFunc func(*net.TCPConn, []byte, int) error
