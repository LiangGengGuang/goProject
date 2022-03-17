package main

import (
	"fmt"
	"time"
)

//子 goroutine
func newTask() {
	i := 0
	for {
		i++
		fmt.Printf("new Goroutine : i = %d\n", i)
		time.Sleep(1 * time.Second)
	}
}

//主 goroutine
func main() {
	go newTask()

	fmt.Println("main Goroutine exit")

	//i := 0
	//for {
	//	i++
	//	fmt.Printf("main Goroutine : i = %d\n", i)
	//	time.Sleep(1 * time.Second)
	//}

}
