package main

import (
	"fmt"
)

func main(){
	//a := 111
	b := "111"
	var arg interface{}
	arg = b
	value, ok := arg.(string)
	if !ok {
		fmt.Println("arg is not string type")
	} else {
		fmt.Println("arg is string type, value = ", value)
	}

}
