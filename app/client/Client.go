package main

import (
	"fmt"
	"io"
	"net"
	"time"

	"github.com/ibuprofen/Tin/tnet"
)

func main() {
	fmt.Println("Client Test ... start")
	// 3秒之后发起测试请求，给服务端开启服务的机会
	time.Sleep(3 * time.Second)

	conn, err := net.Dial("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println("client start err, exit!")
		return
	}

	for {
		dp := tnet.NewDataPack()
		msg, err := dp.Pack(tnet.NewMessage(0, []byte("Zinx V0.5 Client Test Message")))
		if err != nil {
			fmt.Println("pack error msg id = 0")
			return
		}

		_, err = conn.Write(msg)
		if err != nil {
			fmt.Println("write error err ", err)
			return
		}

		dp = tnet.NewDataPack()
		headLen := dp.GetHeadLen()
		buf := make([]byte, headLen)
		if _, err = io.ReadFull(conn, buf); err != nil {
			fmt.Println("read head err ", err)
			break
		}
		msgHead, err := dp.Unpack(buf)
		if err != nil {
			fmt.Println("unpack error msg id = ", msgHead.GetMsgId())
			return
		}
		if msgHead.GetDataLen() > 0 {
			msg := msgHead.(*tnet.Message)
			data := make([]byte, msg.GetDataLen())
			if _, err = io.ReadFull(conn, data); err != nil {
				fmt.Println("read data err ", err)
				return
			}
			fmt.Println("==> Recv Msg: ID=", msg.ID, ", len=", msg.DataLen, ", data=", string(data))
		}

		time.Sleep(1 * time.Second)
	}
}
