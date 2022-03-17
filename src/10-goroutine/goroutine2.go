package main

import (
	"fmt"
	"time"
)

func main() {

	go func() {
		defer fmt.Println("a defer")
		func() {
			defer fmt.Println("b defer")
			//只退出当前func
			//return
			//退出整个func
			//runtime.Goexit()
			fmt.Println("b")
		}()
		fmt.Println("a")
	}()

	for {
		time.Sleep(1 * time.Second)
	}

}
