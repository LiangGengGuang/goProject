package main

import (
	"fmt"
)

var r = 100

//常量
const q = true

func main() {

	//iota 配合const()自增
	const (
		xx = iota
		yy
		zz
	)

	fmt.Println("xx", xx)
	fmt.Println("yy", yy)
	fmt.Println("zz", zz)
	fmt.Println("q", q)
}
