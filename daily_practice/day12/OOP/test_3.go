package main

import "fmt"

type AnimalIF interface {
	Sleep()
	GetColor() string
	GetType() string
}

type Cat struct {
	color string
}

func (this *Cat) Sleep()  {
	fmt.Println("Cat is Sleep !")
}

func (this *Cat) GetColor() string {
	return this.color
}

func (this *Cat) GetType() string {
	return "Cat"
}

type Dog struct {
	color string
}

func (this *Dog) Sleep()  {
	fmt.Println("Dog is Sleep !")
}

func (this *Dog) GetColor() string {
	return this.color
}

func (this *Dog) GetType() string {
	return "Dog"
}

func showAnimal(animal AnimalIF) {
	animal.Sleep()
}

func main() {
/*	var animal AnimalIF
	animal = &Cat{
		color: "green",
	}
	animal.Sleep()

	animal = &Dog{
		color: "blue",
	}
	animal.Sleep() */
	cat := Cat{"green"}
	dog := Dog{"blue"}

	showAnimal(&cat)
	showAnimal(&dog)
}
