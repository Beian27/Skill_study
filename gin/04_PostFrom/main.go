package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.POST("/form_post", func(c *gin.Context) {
		messgae := c.PostForm("message")
		nick := c.DefaultPostForm("nick", "defaultValue")

		c.JSON(200, gin.H{
			"status":  http.StatusOK,
			"message": messgae,
			"nick":    nick,
		})
	})

	r.Run(":8989")
}
