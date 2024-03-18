package main

import (
	"github.com/AMCodeBytes/go-rest-film-review/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// Default returns an engine instance with logger & recovery middleware
	server := gin.Default()

	// Initial the routes
	routes.RegisterRoutes(server)

	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	server.Run()
}
