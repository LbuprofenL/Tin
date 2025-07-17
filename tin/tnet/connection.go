package tnet

import (
	"fmt"
	"net"

	"github.com/ibuprofen/Tin/tin/tinface"
)

type Connection struct {
	// 当前连接的socket TCP套接字
	Conn *net.TCPConn
	// 当前连接的ID 也可以称作为SessionID，ID全局唯一
	ConnID uint32
	// 当前连接的关闭状态
	isClosed bool

	Router tinface.IRouter

	// 告知该链接已经退出/停止的channel
	ExitBuffChan chan bool
}

// NewConntion 创建连接的方法
func NewConntion(conn *net.TCPConn, connID uint32, router tinface.IRouter) *Connection {
	c := &Connection{
		Conn:         conn,
		ConnID:       connID,
		isClosed:     false,
		Router:       router,
		ExitBuffChan: make(chan bool, 1),
	}

	return c
}

// StartReader 读取客户端数据
func (c *Connection) StartReader() {
	fmt.Println("Reader Goroutine is  running")
	defer fmt.Println(c.RemoteAddr().String(), " conn reader exit!")
	defer c.Stop()

	for {
		// 读取客户端数据
		buf := make([]byte, 512)
		_, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("recv buf err ", err)
			c.ExitBuffChan <- true
			continue
		}
		req := Request{
			conn: c,
			data: buf,
		}

		// 调用当前连接所绑定的handleAPI
		go func(req tinface.IRequest) {
			c.Router.PreHandle(req)
			c.Router.Handle(req)
			c.Router.PostHandle(req)
		}(&req)
	}
}

// Start 启动连接
func (c *Connection) Start() {
	go c.StartReader()

	for range c.ExitBuffChan {
		return
	}
}

// Stop 停止连接，结束当前连接状态M
func (c *Connection) Stop() {
	// 1. 如果当前链接已经关闭
	if c.isClosed {
		return
	}
	c.isClosed = true

	// TODO Connection Stop() 如果用户注册了该链接的关闭回调业务，那么在此刻应该显示调用

	// 关闭socket链接
	c.Conn.Close()

	// 通知从缓冲队列读数据的业务，该链接已经关闭
	c.ExitBuffChan <- true

	// 关闭该链接全部管道
	close(c.ExitBuffChan)
}

// GetTCPConnection 从当前连接获取原始的socket TCPConn
func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

// GetConnID 获取当前连接ID
func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

// RemoteAddr 获取远程客户端地址信息
func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}
