package main

import "fmt"

func main() {
	//相当于Java的finally,多个defer执行顺序：先进后出
	defer fmt.Println("main end1")
	defer fmt.Println("main end2")

	fmt.Println("hello 1")
	fmt.Println("hello 2")

	//defer与return的执行顺序
	deferAndReturn()
}

/*
 先return 后defer
*/
func deferAndReturn() bool {
	defer deferMeth()
	return returnMeth()

}
func deferMeth() bool {
	fmt.Println("deferMeth ...")
	return true
}
func returnMeth() bool {
	fmt.Println("returnMeth ...")
	return true
}
