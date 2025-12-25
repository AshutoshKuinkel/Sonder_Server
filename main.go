package main

import (
	"fmt"

	"sonder_Server/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	//db connections

	//middlewares
	r.Use(middlewares.ErrorHandler())
	// routes

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Server Live",
		})
	})

	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{
			"message": fmt.Sprintf("Cannot %s @ %s", c.Request.Method, c.Request.URL.Path),
		})
	})

	r.Run()
}
