package znet

import (
	"fmt"
	"sync"
	"zinx/ziface"
)

/*
	ConnManager 链接管理具体实现
*/
type ConnManager struct {

	//链接字典
	ConnMap map[uint32]ziface.IConnection

	//链接字典的读写锁
	Lock sync.RWMutex
}

// NewConnManager 初始化链接容器方法
func NewConnManager() ziface.IConnManager {
	return &ConnManager{
		ConnMap: make(map[uint32]ziface.IConnection),
	}
}

func (cm *ConnManager) Add(conn ziface.IConnection) {

	cm.Lock.Lock()
	defer cm.Lock.Unlock()
	cm.ConnMap[conn.GetConnID()] = conn

	fmt.Println("add connection,connId=", conn.GetConnID())
}

func (cm *ConnManager) Remove(connId uint32) {

	cm.Lock.Lock()
	defer cm.Lock.Unlock()
	delete(cm.ConnMap, connId)

	fmt.Println("remove connection,connId=", connId)
}

func (cm *ConnManager) Get(connId uint32) ziface.IConnection {

	cm.Lock.RLock()
	defer cm.Lock.RUnlock()
	conn := cm.ConnMap[connId]
	fmt.Println("remove connection,connId=", connId)
	return conn
}

func (cm *ConnManager) Quantity() int {

	//获取当前链接的数量
	cm.Lock.RLock()
	defer cm.Lock.RUnlock()
	return len(cm.ConnMap)
}

func (cm *ConnManager) ClearAll() {

	cm.Lock.Lock()
	defer cm.Lock.Unlock()

	//停止并删除全部的连接信息
	for connID, conn := range cm.ConnMap {

		//停止
		conn.Stop()

		//删除
		delete(cm.ConnMap, connID)
	}

	fmt.Println("clear all connection")
}

func (cm *ConnManager) ClearOne(ConnId uint32) {

	cm.Lock.Lock()
	defer cm.Lock.Unlock()

	if conn := cm.Get(ConnId); conn != nil {

		//停止
		conn.Stop()

		//删除
		delete(cm.ConnMap, ConnId)
	}
	fmt.Println("Clear one connection")
}
