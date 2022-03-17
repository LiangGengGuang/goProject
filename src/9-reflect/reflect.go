package main

//反射
import (
	"encoding/json"
	"fmt"
	"reflect"
)

type User struct {
	Name string `info:"Name" doc:"姓名" json:"Name"` //json => 转json格式标签
	Sex  string `info:"Sex" doc:"性别" json:"Sex"`
	Age  int    `info:"Age" doc:"年龄" json:"Age"`
}

func (this User) Call() {
	fmt.Println("user is called ...")
	fmt.Printf("%v\n", this)
}

func reflectMethod(input interface{}) {

	//获取类型
	typeInput := reflect.TypeOf(input)
	fmt.Println("typeInput", typeInput.Name())
	//获取值
	valInput := reflect.ValueOf(input)
	fmt.Println("valInput", valInput)
	//获取User内部参数的类型和值
	for i := 0; i < typeInput.NumField(); i++ {
		file := typeInput.Field(i)
		val := valInput.Field(i).Interface()
		fmt.Printf("%s:%v = %v\n", file.Name, file.Type, val)
	}
	//通过typeOf获取User的方法，并调用
	for i := 0; i < typeInput.NumMethod(); i++ {
		m := typeInput.Method(i)
		fmt.Printf("%s:%v\n", m.Name, m.Type)
	}
}

func findTag(input interface{}) {
	t := reflect.TypeOf(input).Elem()
	for i := 0; i < t.NumField(); i++ {
		info := t.Field(i).Tag.Get("info")
		doc := t.Field(i).Tag.Get("doc")
		fmt.Println("info", info)
		fmt.Println("doc", doc)
	}
}

func main() {
	user := User{"lgg", "male", 28}

	reflectMethod(user)

	findTag(&user)

	//结构体转为json
	jsonStr, err := json.Marshal(user)
	if err != nil {
		fmt.Println("error")
		return
	}
	//Printf格式化输出
	fmt.Printf("user json = %s\n", jsonStr)

	newJsonStr := User{}

	err = json.Unmarshal(jsonStr, &newJsonStr)
	if err != nil {
		fmt.Println("error")
		return
	}
	fmt.Println("newUser", newJsonStr)

}
