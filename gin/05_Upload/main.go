package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	r := gin.Default()

	//r.MaxMultipartMemory = 64 << 20
	r.POST("/upload", func(c *gin.Context) {
		// 单文件
		file, err := c.FormFile("file")
		if err != nil {
			log.Fatal(err)
			return
		}
		log.Println(file.Filename, file.Header, file.Size)

		// 上传文件到指定路径
		//c.SaveUploadedFile(file, "/Users/mac/GolangProjects/Skill_study/skill/05_Upload/file")

		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})
	r.Run(":8989")
}
