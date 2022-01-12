package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default() // 携带中间件启动  Logger
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "get",
		})
	})

	r.POST("/post", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "post",
		})
	})

	r.PUT("/put", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "put",
		})
	})

	r.DELETE("/delete", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "del",
		})
	})

	r.PATCH("/patch", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "patch",
		})
	})

	r.HEAD("/head", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "head",
		})
	})

	r.OPTIONS("/options", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"messgae": "options",
		})
	})
	r.Run(":8989")
}
