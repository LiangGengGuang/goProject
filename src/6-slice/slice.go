package main

import "fmt"

func main() {
	//开辟空间并赋值
	slice1 := []int{1, 2, 3}
	fmt.Println("slice1：", slice1)

	//动态数组，开辟空间，长度为2，容量为5，默认值为0
	//var slice2 []int
	//slice2 = make([]int, 2, 5)
	//var slice2 []int = make([]int, 2, 5)
	slice2 := make([]int, 2, 5) //cap参数未填写，默认与len参数相同
	fmt.Printf("slice2：%d,len：%d,cap：%d,", slice2, len(slice2), cap(slice2))
	fmt.Println()

	//切片内容追加,当数值长度超过原有容量后，切片会在原容量基础上的自动扩容一倍
	slice2 = append(slice2, 1)
	slice2 = append(slice2, 2)
	fmt.Printf("slice2：%d,len：%d,cap：%d,", slice2, len(slice2), cap(slice2))
	fmt.Println()

	var s3 []int
	//截取，参数一：截取起点，参数二：截取长度
	s2 := slice1[0:2]
	fmt.Println("s2：", s2)
	copy(s3, s2)
	fmt.Println("s3:", s3)
}
