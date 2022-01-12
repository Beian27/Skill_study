package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/welcome", func(context *gin.Context) {
		firstname := context.DefaultQuery("firstname", "guest")
		lastname := context.Query("lastname")

		context.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	})

	r.Run(":8989")
}
