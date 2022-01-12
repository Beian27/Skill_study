package main

import(
	"bufio"
	"fmt"
	"io"
	"os"
)

func main(){
	fp, err := os.Open("/Users/mac/GolandProjects/Skill_study/day15/a.txt")
	if err != nil {
		fmt.Println("文件打开失败！")
		return
	}

	defer fp.Close()

	reader := bufio.NewReader(fp)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Print(line)
	}

}
