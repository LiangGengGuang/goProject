package main

import "fmt"

//继承
type Human struct {
	Name   string
	gender string
}

type superMan struct {
	Human
	level int
}

func (this Human) Eat() {
	fmt.Println("Human Eat ...")

}

func (this Human) Walk() {
	fmt.Println("Human Walk ...")
}

func (this superMan) Eat() {
	fmt.Println("superMan Walk ...")
}

func (this superMan) Fly() {
	fmt.Println("superMan Fly ...")
}

func main() {

	h := Human{"张三", "female"}
	fmt.Println("human:", h)
	h.Eat()
	h.Walk()

	//s := superMan{Human{"李四", "female"}, 80}
	var s superMan
	s.Name = "李四"
	s.gender = "female"
	s.level = 88

	fmt.Println("superMan:", s)
	s.Eat()
	s.Walk()
	s.Fly()
}
