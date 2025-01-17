package main

import (
	"deployapi/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create a Gin router
	r := gin.Default()

	// Enable CORS for frontend-backend communication
	r.Use(corsMiddleware())

	v1 := r.Group("/api/v1")
	// Define the API endpoint
	v1.POST("/process", controller.Handler)
	v1.GET("/m/:msg", controller.Hello)

	// Start the server on port 8080
	r.Run(":8080")
}

// CORS middleware to handle cross-origin requests
func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	}
}
