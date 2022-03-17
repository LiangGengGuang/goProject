package main

import "fmt"

//声明一种新的数据类型myInt,是int的别名
type myInt int

//封装
//定义一个结构体
type Book struct {
	Title string
	Auth  string
	Price float64
}

//指针传递
func changBook(book *Book) {
	book.Price = 62.8
}

func (this Book) getTitle() {
	fmt.Println("Title= ", this.Title)
}

func (this *Book) setPrice(newPrice float64) {
	this.Price = newPrice
}

func main() {

	var book1 Book
	book1.Title = "三国演义"
	book1.Auth = "罗贯中"
	book1.Price = 89.8

	fmt.Printf("打折前：%v\n", book1)
	changBook(&book1)
	fmt.Printf("7折后：%v\n", book1)
	book1.setPrice(44.9)
	fmt.Printf("5折后：%v\n", book1)
}
