package main

import "fmt"

//关闭缓冲
func main() {

	c := make(chan int)

	//与main并行，需要通过channel将func中的值取出，在main中使用
	go func() {

		for i := 0; i < 4; i++ {
			c <- i
		}
		//关闭channel
		close(c)
	}()
	/*
		for {
			//第二ok为true表示channel未关闭，为false表示channel关闭
			if numb, ok := <-c; ok {
				fmt.Println("num = ", numb)
			} else {
				break
			}
		}
	*/

	//使用ranget迭代拿去channel中的数据
	for numb := range c {
		fmt.Println("num = ", numb)
	}
}
