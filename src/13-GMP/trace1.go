package main

import (
	"fmt"
	"os"
	"runtime/trace"
)

/*
	可视化调试
	1.go run trace1.go
	2.go tool trace trace1.out
*/
func main() {

	//启动trace文件
	create, err := os.Create("trace1.out")
	defer create.Close()

	if err != nil {
		panic(err)
	}

	//启动trace
	err = trace.Start(create)
	if err != nil {
		panic(err)
	}

	fmt.Println("hello world")

	//停止trace
	trace.Stop()
}
