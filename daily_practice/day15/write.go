package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("/Users/mac/GolandProjects/Skill_study/day15/a.txt", os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	bytes := []byte("您好！！！！！张宏宇！！！")
	content, err := file.Write(bytes)
	 if  err != nil {
		 fmt.Println(err)
	 }
	 fmt.Println(content)
	 fmt.Println("写入成功！！！")
}
