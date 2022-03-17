package main

import "fmt"

//select监控多路channel

func fibonacii(c, quit chan int) {
	x, y := 1, 1

	for {
		select {
		case c <- x:
			x = y
			y = x + y
		case <-quit:
			fmt.Println("quit")
			return
		}

	}

}
func main() {

	c := make(chan int)
	quit := make(chan int)

	//与main并行，需要通过channel将func中的值取出，在main中使用
	go func() {

		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()

	fibonacii(c, quit)

  }
