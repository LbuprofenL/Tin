package utils

import (
	"encoding/json"
	"os"

	"github.com/ibuprofen/Tin/tinface"
)

type GlobalObj struct {
	Server  *tinface.IServer
	Host    string
	TCPPort int
	Name    string
	Version string // 当前tin版本号

	MaxConn        int    // 当前服务器主机允许的最大链接数
	MaxPackageSize uint32 // 当前服务器主机允许的最大数据包字节数
}

var GlobalObject *GlobalObj

func (g *GlobalObj) Reload() {
	data, err := os.ReadFile("conf/tin.json")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(data, g)
	if err != nil {
		panic(err)
	}
}

func init() {
	// 初始化GlobalObject变量，设置一些默认值
	GlobalObject = &GlobalObj{
		Name:           "TinServerApp",
		Version:        "V0.4",
		TCPPort:        7777,
		Host:           "0.0.0.0",
		MaxConn:        12000,
		MaxPackageSize: 4096,
	}

	// 从配置文件中加载一些用户配置的参数
	GlobalObject.Reload()
}
