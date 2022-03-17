package main

import "fmt"

func main() {
	method("hello", "world")

	fmt.Println("method2=>", method2(1))

	res1, res2 := method3(1)
	fmt.Println("method3=> res1=", res1, "res2=", res2)

	res1, res2 = method4()
	fmt.Println("method4=> res1=", res1, "res2=", res2)

	res1, res2 = method5()
	fmt.Println("method5=> res1=", res1, "res2=", res2)
}

func method(x string, y string) {
	fmt.Println("method=> x+y=", x+y)
}

func method2(x int) int {
	return x * 2
}

func method3(x int) (int, int) {
	return x, x * 2
}

func method4() (r1 int, r2 int) {
	//无赋值，默认为0
	return r1, r2
}
func method5() (r1, r2 int) {
	r1 = 10
	r2 = 20
	return r1, r2
}
