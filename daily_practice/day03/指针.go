package main

import "fmt"

func main() {
	var a int = 20
	var ip *int
	fmt.Printf("ip 的值为 : %x\n", ip  )
	if ip == nil {
		fmt.Printf("true\n")
	}
	// &a 取a变量的地址
	ip = &a 
	fmt.Printf("a 变量的地址：%x\n", &a)
	fmt.Printf("ip 变量储存的指针地址: %x\n", ip )
	fmt.Printf("*ip 变量的值: %d\n", *ip )
}