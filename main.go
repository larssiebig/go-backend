package main

import (
	"go-backend/handlers"
	"go-backend/middleware"
	"go-backend/utils"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize DB connection
	utils.InitDB()

	// Initialize Gin router
	r := gin.Default()

	// Auth routes
	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)
	r.DELETE("/user/:id", handlers.DeleteUser)

	// Protected routes (JWT required)
	protected := r.Group("/api")
	protected.Use(middleware.JWTAuth())
	protected.GET("/profile", func(c *gin.Context) {
		userID := c.MustGet("userID")
		c.JSON(200, gin.H{"userID": userID, "message": "Protected route"})
	})

	// Start the server
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
