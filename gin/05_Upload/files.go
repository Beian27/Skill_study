package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.Default()

	r.POST("/upload", func(c *gin.Context) {
		form, _ := c.MultipartForm()
		files := form.File["upload"]

		for _, file := range files {
			log.Println(file.Filename)
			c.SaveUploadedFile(file, "skill/05_Upload/file/"+file.Filename)
			c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%S", file.Filename))
			c.File("skill/05_Upload/file/" + file.Filename)
		}
		//c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
	})
	r.Run(":8989")
}
