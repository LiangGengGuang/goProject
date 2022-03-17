package lib1

import "fmt"

//方法名首字母大写表示可以被外部调用
func LibMethod() {
	fmt.Println("lib1Method ...")
}

func init() {
	fmt.Println("lib1 init() ...")
}
