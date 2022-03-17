package main

import "fmt"

//多态 必须实现接口所有方法，才表示实现该接口
//本质是指针
type Animal interface {
	Sound()
	Eat()
}

//实体对象
type Cat struct {
}

func (this *Cat) Sound() {
	fmt.Println("喵喵猫...")

}
func (this *Cat) Eat() {
	fmt.Println("猫吃鱼")
}

type Dog struct {
}

//实体对象
func (this Dog) Sound() {
	fmt.Println("汪汪汪...")

}
func (this Dog) Eat() {
	fmt.Println("狗吃骨头")
}

func ShowAnimal(animal Animal) {
	animal.Sound()
	animal.Eat()

}

func main() {

	var cat Animal
	cat = &Cat{}
	ShowAnimal(cat)

	dog := Dog{}
	ShowAnimal(&dog)
}
