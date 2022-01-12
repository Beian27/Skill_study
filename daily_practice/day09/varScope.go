package main

import "fmt"

var global int

func main() {

	var local1, local2 int

	local1 = 8

	local2 = 10	
	
	global = local1 - local2
	
	fmt.Println(global)

	global = local1 * local2

	fmt.Println(global)
}
