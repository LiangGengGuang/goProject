package main

/*
import "fmt"
import "time"
*/
import (
	"fmt"
	"time"
)

var (
	o int
	p string = ""
)
var r = 100

func main() {

	x := 1 //这种不带声明格式的只能在函数体中出现
	var y, z = 2, 3
	fmt.Println(x == y)
	fmt.Println("y:", y)
	fmt.Println("z:", z)

	var i int      //默认0
	var j1 float32 //默认0
	var j2 float64 //默认0
	fmt.Println("i:", i)
	fmt.Println("j1:", j1)
	fmt.Println("j2:", j2)

	time.Sleep(5 * time.Second)

	g, h := 123, "hello" //可同时给变量设置不同类型的值
	fmt.Println("g:", g)
	fmt.Println("h:", h)

	var flag bool //默认false
	fmt.Println("flag:", flag)
}
