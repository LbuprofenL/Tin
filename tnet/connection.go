package tnet

import (
	"errors"
	"fmt"
	"io"
	"net"

	"github.com/ibuprofen/Tin/tinface"
)

type Connection struct {
	// 当前连接的socket TCP套接字
	Conn *net.TCPConn
	// 当前连接的ID 也可以称作为SessionID，ID全局唯一
	ConnID uint32
	// 当前连接的关闭状态
	isClosed bool

	Handler tinface.IMsgHandler

	// 告知该链接已经退出/停止的channel
	ExitBuffChan chan bool
}

// NewConntion 创建连接的方法
func NewConntion(conn *net.TCPConn, connID uint32, handler tinface.IMsgHandler) *Connection {
	c := &Connection{
		Conn:         conn,
		ConnID:       connID,
		isClosed:     false,
		Handler:      handler,
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
		// 拆包
		dp := &DataPack{}

		headLen := dp.GetHeadLen()
		buf := make([]byte, headLen)
		if _, err := io.ReadFull(c.GetTCPConnection(), buf); err != nil {
			fmt.Println("read head err ", err)
			continue
		}
		// 拆包, 得到msgID和dataLen
		msg, err := dp.Unpack(buf)
		if err != nil {
			fmt.Println("unpack err ", err)
			continue
		}
		// 根据dataLen再读取data
		if msg.GetDataLen() > 0 {
			data := make([]byte, msg.GetDataLen())
			if _, err := io.ReadFull(c.GetTCPConnection(), data); err != nil {
				fmt.Println("read data err ", err)
				continue
			}
			msg.SetData(data)
		}
		req := Request{
			conn: c,
			msg:  msg,
		}

		// 调用当前连接所绑定的handleAPI
		go c.Handler.DoMsgHandler(&req)
	}
}

// SendMsg 发送消息给客户端
func (c *Connection) SendMsg(msgID uint32, data []byte) error {
	if c.isClosed {
		return errors.New("connection is closed")
	}

	dp := &DataPack{}
	msg, err := dp.Pack(NewMessage(msgID, data))
	if err != nil {
		fmt.Println("Pack error msg id = ", msgID)
		return errors.New("Pack error msg ")
	}
	_, err = c.Conn.Write(msg)
	if err != nil {
		fmt.Println("send msg err ", err)
		c.ExitBuffChan <- true
		return errors.New("conn Write error")
	}
	fmt.Printf("send msg to client[%s] success, msgID: %d, dataLen: %d\n", c.RemoteAddr().String(), msgID, len(data))
	return nil
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
