package main

import (
	"github.com/gin-gonic/gin"
)

type Person struct {
	Name     string `json:"name" binding:"required"`
	Address  string `json:"address"`
	Birthday string `json:"birthday"`
}

func main() {
	r := gin.Default()

	r.POST("/testbind", startPage)

	r.Run(":8989")
}

func startPage(c *gin.Context) {
	var person Person
	err := c.ShouldBindJSON(&person)
	if err != nil {
		c.JSON(200, gin.H{
			"data":    gin.H{},
			"message": "出错了！",
		})
	} else {
		c.JSON(200, gin.H{
			"data":    person,
			"message": "成功了！",
		})
	}

}
