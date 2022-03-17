package main

import (
	//"4-init/lib1"
	//"4-init/lib2"
	//"4-init/lib3"
	_ "4-init/lib1"      //匿名别名，只能执行当前包内的init方法
	. "4-init/lib2"      //包内方法直接引用 不建议使用
	myLib3 "4-init/lib3" //自定义别名
	"fmt"
)

func main() {
	//引用外部包的调用顺序
	//lib1.LibMethod()
	LibMethod()
	myLib3.LibMethod()
}

func init() {
	fmt.Println("main init() ...")
}
