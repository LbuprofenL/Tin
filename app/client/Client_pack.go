package main

// import (
// 	"fmt"
// 	"net"

// 	"github.com/ibuprofen/Tin/tnet"
// )

// func main() {
// 	// 客户端goroutine，负责模拟粘包的数据，然后进行发送
// 	conn, err := net.Dial("tcp", "127.0.0.1:7777")
// 	if err != nil {
// 		fmt.Println("client dial err:", err)
// 		return
// 	}

// 	// 创建一个封包对象 dp
// 	dp := tnet.NewDataPack()

// 	// 封装一个msg1包
// 	msg1 := &tnet.Message{
// 		ID:      0,
// 		DataLen: 5,
// 		Data:    []byte{'h', 'e', 'l', 'l', 'o'},
// 	}

// 	sendData1, err := dp.Pack(msg1)
// 	if err != nil {
// 		fmt.Println("client pack msg1 err:", err)
// 		return
// 	}

// 	msg2 := &tnet.Message{
// 		ID:      1,
// 		DataLen: 7,
// 		Data:    []byte{'w', 'o', 'r', 'l', 'd', '!', '!'},
// 	}
// 	sendData2, err := dp.Pack(msg2)
// 	if err != nil {
// 		fmt.Println("client temp msg2 err:", err)
// 		return
// 	}

// 	// 将sendData1，和 sendData2 拼接一起，组成粘包
// 	sendData1 = append(sendData1, sendData2...)

// 	// 向服务器端写数据
// 	_, err = conn.Write(sendData1)
// 	if err != nil {
// 		fmt.Println("client pack msg1 err:", err)
// 		return
// 	}
// 	fmt.Println("send msg1 to server success")
// 	// 关闭连接
// 	conn.Close()
// }
