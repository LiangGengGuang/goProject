package main

import (
	"fmt"
	"time"
)

//有缓冲
func main() {

	c := make(chan int, 3) //带有缓存的channel，缓冲区大小为3

	fmt.Println("len(c)= ", len(c), "，cap(c)= ", cap(c))

	//与main并行，需要通过channel将func中的值取出，在main中使用
	go func() {
		defer fmt.Println("子goroutine结束")

		for i := 0; i < 4; i++ {
			c <- i
			fmt.Println("goroutine 运行... len(c)= ", len(c), "，cap(c)= ", cap(c))
		}
	}()

	time.Sleep(2 * time.Second)
	for i := 0; i < 3; i++ {
		num := <-c
		fmt.Println("num = ", num)
	}
	fmt.Println("main goroutine 结束")
}
