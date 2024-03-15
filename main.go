package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Default returns an engine instance with logger & recovery middleware
	r := gin.Default()

	// Setup a route for /ping
	r.GET("/ping", func(c *gin.Context) {
		// Send a response of pong
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	r.Run()
}
