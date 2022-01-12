package main

import "fmt"

type Hero struct {
	Name 	string
	Ad 		int
	Level 	int
}

func (this *Hero) GetName() string {
	return this.Name
}

func (this *Hero) SetName(newName string) {
	this.Name = newName
}

func (this *Hero) Show() {
	fmt.Println(this)
}

func main() {
	hero := Hero{Name: "zhy", Ad: 100, Level: 1}
	fmt.Println(hero.GetName())
	hero.SetName("张宏宇")
	hero.Show()
	a := 1
	b := &a
	fmt.Println(&a)
	fmt.Println(*b)
}
