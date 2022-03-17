package lib3

import "fmt"

//方法名首字母大写表示可以被外部调用
func LibMethod() {
	fmt.Println("lib3Method ...")
}

func init() {
	fmt.Println("lib3 init() ...")
}
