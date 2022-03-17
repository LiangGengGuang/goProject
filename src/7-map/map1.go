package main

import "fmt"

func main() {
	var myMap map[int]string
	if myMap == nil {
		fmt.Println("myMap为空")
	}

	myMap = make(map[int]string, 10)
	myMap[1] = "java"
	myMap[2] = "C"
	myMap[3] = "go"
	myMap[4] = "python"
	fmt.Println("myMap", myMap)

	myMap2 := make(map[int]string)
	myMap2[1] = "java"
	myMap2[2] = "C"
	myMap2[3] = "go"
	myMap2[4] = "python"
	fmt.Println("myMap2", myMap2)

	myMap3 := map[int]string{
		1: "java",
		2: "C",
		3: "go",
		4: "python",
	}
	fmt.Println("myMap3", myMap3)
}
