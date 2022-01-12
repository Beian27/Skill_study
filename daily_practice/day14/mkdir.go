package main

import (
	"fmt"
	"os"
)

func main(){
	err := os.Mkdir("/Users/mac/GolandProjects/Skill_study/day15", 777)
	if err != nil {
		fmt.Println(err)
	}

}
