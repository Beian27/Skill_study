package main

import "fmt"

type Rectangle struct {
	width, height float64
}

func area(r Rectangle) float64 {
	return r.width * r.height
}

func (r Rectangle) area_1() float64 {
	return r.width * r.height
}

func main() {
	r1 := Rectangle{12, 2}
	r2 := Rectangle{12, 12}
	r3 := Rectangle{12.5, 120}

	fmt.Println("Area of r1 is:", area(r1))
	fmt.Println("Area of r2 is:", area(r2))
	fmt.Println("Area of r2 is:", r3.area_1())

}
