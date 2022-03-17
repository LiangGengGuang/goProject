package main

import "fmt"

//普通数组
func main() {

	var array1 = [3]int{1, 2, 3}
	fmt.Println(array1)
	for i := 0; i < len(array1); i++ {
		fmt.Println(array1[i])
	}

	var array2 [3]int
	array2[0] = 7
	array2[1] = 8
	array2[2] = 9
	fmt.Println("array3长度:", len(array2))
	fmt.Printf("arrays type = %T", array2)
	for index, value := range array2 {
		fmt.Println("index:", index, "value", value)
	}

}
