package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	filePath := "/Users/mac/GolandProjects/Skill_study/day15/a.txt"
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", string(content))
}
