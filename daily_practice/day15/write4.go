package main

import (
	"fmt"
	"os"
)

func main() {
	create, err := os.Create("/Users/mac/GolandProjects/Skill_study/day15/b.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer create.Close()

	for i := 0; i < 1000; i++ {
		content := fmt.Sprintf("%s:%d\r\n", "Hello Go", i)
		create.WriteString(content)
		create.Write([]byte("I love go\r\n"))
	}


}
