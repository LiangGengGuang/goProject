package main

import "fmt"

//切片sile 动态数组
func main() {

	array := []int{4, 5, 6, 7}
	printArray(array)

	fmt.Println("====")
	for _, value := range array {
		fmt.Println("value", value)
	}
}

//切片是应用传递
func printArray(array []int) {

	for _, value := range array {
		fmt.Println("value", value)
	}
	array[0] = 8
}
