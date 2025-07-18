package tnet

import (
	"errors"
	"fmt"
	"net"
	"time"

	"github.com/ibuprofen/Tin/tinface"
	"github.com/ibuprofen/Tin/utils"
)

// iServer 接口实现，定义一个Server服务类
type Server struct {
	Name      string
	IPVersion string
	IP        string
	Port      int
	Router    tinface.IRouter
}

// CallBackToClient 定义当前客户端链接的handle api
func CallBackToClient(conn *net.TCPConn, data []byte, cnt int) error {
	// 回显业务
	fmt.Println("[Conn Handle] CallBackToClient ... ")
	if _, err := conn.Write(data[:cnt]); err != nil {
		fmt.Println("write back buf err ", err)
		return errors.New("CallBackToClient error")
	}
	return nil
}

// ============== 实现tinface.IServer 里的全部接口方法 ===========

func (s *Server) Start() {
	fmt.Println("[START] Server is starting...")
	fmt.Printf("[Tin] Version: %s, MaxConn: %d,  MaxPackageSize: %d\n",
		utils.GlobalObject.Version,
		utils.GlobalObject.MaxConn,
		utils.GlobalObject.MaxPackageSize)
	// 开启一个go去做服务端Linster业务
	go func() {
		// 1 获取一个TCP的Addr
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolve tcp addr err: ", err)
			return
		}

		// 2 监听服务器地址
		listenner, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("listen", s.IPVersion, "err", err)
			return
		}

		// 已经监听成功
		fmt.Println("start Tin server  ", s.Name, " succ, now listenning...")
		var cid uint32
		// 3 启动server网络连接业务
		for {
			// 3.1 阻塞等待客户端建立连接请求
			conn, err := listenner.AcceptTCP()
			if err != nil {
				fmt.Println("Accept err ", err)
				continue
			}

			// 3.2 TODO Server.Start() 设置服务器最大连接控制,如果超过最大连接，那么则关闭此新的连接

			// 3.3 TODO Server.Start() 处理该新连接请求的 业务 方法， 此时应该有 handler 和 conn是绑定的
			dealConn := NewConntion(conn, cid, s.Router)
			cid++
			go dealConn.Start()
		}
	}()
}

func (s *Server) Stop() {
	fmt.Println("[STOP] Tin server , name ", s.Name)

	// TODO  Server.Stop() 将其他需要清理的连接信息或者其他信息 也要一并停止或者清理
}

func (s *Server) Serve() {
	s.Start()

	// TODO Server.Serve() 是否在启动服务的时候 还要处理其他的事情呢 可以在这里添加

	// 阻塞,否则主Go退出， listenner的go将会退出
	for {
		time.Sleep(10 * time.Second)
	}
}

func (s *Server) AddRouter(router tinface.IRouter) {
	s.Router = router
}

// NewServer 创建一个服务器句柄
func NewServer() tinface.IServer {
	utils.GlobalObject.Reload()

	s := &Server{
		Name:      utils.GlobalObject.Name,
		IPVersion: "tcp4",
		IP:        utils.GlobalObject.Host,
		Port:      utils.GlobalObject.TCPPort,
		Router:    nil,
	}

	return s
}
