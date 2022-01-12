package main

import "fmt"

type Human struct {
	name 	string
	sex		string
}

func (h *Human) Eat()  {
	fmt.Println("Human.Eat()")
}

func (h *Human) Walk() {
	fmt.Println("Human.Walk()")
}

type SuperMan struct {
	Human
	level int
}

func (h *SuperMan) Eat()  {
	fmt.Println("SuperMan.Eat()")
}

func (h *SuperMan) Walk() {
	fmt.Println("SuperMan.Walk()")
}


func main() {
	s := SuperMan{Human{name: "张宏宇", sex: "男"}, 1}
	s.Walk()
	s.Eat()
}
