package main

import "fmt"

//无缓冲
func main() {

	c := make(chan int)

	//与main并行，需要通过channel将func中的值取出，在main中使用
	go func() {
		defer fmt.Println("goroutine结束")
		fmt.Println("goroutine 运行...")
		c <- 666
	}()

	num := <-c
	fmt.Println("num = ", num)
	fmt.Println("main goroutine 结束")
}
