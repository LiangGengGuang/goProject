package main

import (
	"fmt"
	"time"
)

/*
	GODEBUG调试
	1.go build trace2.go
	2.GODEBUG=schedtrace=1000 ./trace2 //每1秒打印一次

	例子：	SCHED 0ms: gomaxprocs=12 idleprocs=9 threads=5 spinningthreads=2 idlethreads=1 runqueue=0 [0 0 0 0 0 0 0 0 0 0 0 0]
			hello GMP
			SCHED 1008ms: gomaxprocs=12 idleprocs=12 threads=5 spinningthreads=0 idlethreads=3 runqueue=0 [0 0 0 0 0 0 0 0 0 0 0 0]

	gomaxprocs：配置的处理器数P
	idleprocs：空闲的P数量
	threads：开辟M的数量，运行期管理的线程数
	spinningthreads：自旋线程
	idlethreads：空闲M的数量,
	runqueue：全局队列中G的数量
	[0,0,...]：	本地run队列中G的数量
*/

func main() {

	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
		fmt.Println("hello GMP")
	}
}
