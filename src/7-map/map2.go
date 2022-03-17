package main

import "fmt"

func main() {

	city := map[string]string{
		"China": "Beijing",
		"Japan": "Toyo",
		"USA":   "NewYork",
	}
	printMap(city)
	fmt.Println("============")
	//删除
	delete(city, "Japan")
	//修改
	city["USA"] = "LA"
	printMap(city)
}

func printMap(city map[string]string) {
	//引用传递
	for key, val := range city {
		fmt.Println("key=", key, "val=", val)
	}
}
