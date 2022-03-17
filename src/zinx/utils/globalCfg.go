package utils

import (
	"encoding/json"
	"io/ioutil"
	"zinx/ziface"
)

/*
	存储一切有关Zinx框架的全局配置，供其他模块使用
	一些参数通过zinx.json由用户自定义配置
*/
type GlobalObj struct {
	TcpServer      ziface.IServer //当前zinx全局的server对象
	Host           string         //当前zinx监听的IP对象
	TcpPort        int            //当前zinx监听的端口对象
	Name           string         //当前zinx服务的名称
	Version        string         //当前zinx服务的版本号
	MaxCon         int            //当前zinx服务的最大链接数
	MaxPackageSize uint32         //当前zinx服务数据库最大值
}

/*
	定义一个全局的对外GlobalObject
*/
var GlobalObject *GlobalObj

/*
	zinx.json读取自定义参数
*/
func (c *GlobalObj) Reload() {
	file, err := ioutil.ReadFile("conf/zinx.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(file, &GlobalObject)
	if err != nil {
		panic(err)
	}
}

//读取用户配置好的zinx.json文件
func init() {
	GlobalObject := &GlobalObj{
		Host:           "0.0.0.0",
		TcpPort:        8889,
		Name:           "Zinx",
		Version:        "0.4",
		MaxCon:         1000,
		MaxPackageSize: 4096,
	}
	//调用读取外置配置文件
	GlobalObject.Reload()
}
