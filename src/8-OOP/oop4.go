package main

import "fmt"

//interface是万能数据类型
func MyFunc(arg interface{}) {

	fmt.Println(arg)

	//当arg为空接口时，interface会获得"类型断言"的机制，区分数据类型
	value, ok := arg.(string)
	if ok {
		fmt.Println("value", value)
	} else {
		fmt.Println("arg not string")
	}
}

type Car struct {
	brand string
}

func main() {

	car := Car{"BMW"}
	MyFunc(car)
	MyFunc(1)
	MyFunc(1.5)
	MyFunc("lgg")
}
