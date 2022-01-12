package main

import (
	"fmt"
	"strconv"
)

func main(){
	str := "1"
	in, err := strconv.Atoi(str)
	i, err := strconv.ParseInt(str,10,64)
	itoa := strconv.Itoa(1)
	formatInt := strconv.FormatInt(1, 10)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(i)
	fmt.Println(in)
	fmt.Println(itoa)
	fmt.Println(formatInt)
}
