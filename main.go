package main

import (
	"ginapp/pkg/sum"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		ret := sum.Add(1, 2)
		c.JSON(200, gin.H{
			"message": "pong",
			"ret":     ret,
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
