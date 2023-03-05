package main

import "fmt"

/*
扩容机制
	1.在原容量扩大两倍还要小于扩容后的容量时，预估容量就是扩容后的
	2.当数值长度超过原有容量后，切片会在原容量基础上的自动扩容一倍
	3.扩容后容量超过1024之后，转为每次扩容1/4左右
	4.若append多个元素，且double后的容量不能容纳，直接使用预估的容量；
	  扩容公式：预估容量*元素类型字节数（如果相乘结果不在内存对齐上，则向上取值 ）/元素类型字节数=最终容量
*/
func main() {

	a := []byte{1, 0}
	a = append(a, 1, 1, 1)
	fmt.Println("cap of a is ", cap(a))

	b := []int{23, 51}
	b = append(b, 4, 5, 6)
	fmt.Println("cap of b is ", cap(b))

	//1.12版本之后对int32进行改造，原扩容后cap为8
	c := []int32{1, 23}
	c = append(c, 2, 5, 6)
	fmt.Println("cap of c is ", cap(c))

	type D struct {
		age  byte
		name string
	}
	d := []D{
		{1, "123"},
		{2, "234"},
	}
	d = append(d, D{4, "456"}, D{5, "567"}, D{6, "678"})
	fmt.Println("cap of d is ", cap(d))

	e := []int32{1, 2}
	e = append(e, 3)
	fmt.Println("cap of c is ", cap(e))
}
