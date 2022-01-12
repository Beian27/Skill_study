package main

import (
	"fmt"
	"os"
	"runtime/trace"
)

func main(){

	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	err = trace.Start(f)
	if err != nil {
		panic(err)
	}

	fmt.Println("Hello GMP")

	trace.Stop()

}
