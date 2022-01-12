package main

import (
	"fmt"
	"strings"
)

func main(){
	//str := "hello,world!!"
	arr := strings.Join([]string{"hello","world"}, ",")
	fmt.Println(arr)
	fmt.Printf("%T\n", arr)
	fmt.Println(f(1))
	a := g()
	fmt.Println(&a)
}

func f(x int) *int {
	return &x
}

func g() int{
	x := new(int)
	return *x
}
