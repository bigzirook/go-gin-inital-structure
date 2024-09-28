package main

import (
	"github.com/bigzirook/movie-ticket-booking/config"
	"github.com/bigzirook/movie-ticket-booking/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Initialize database
	config.ConnectDB()

	// Register routes
	routes.RegisterRoutes(router)

	// Start the server
	router.Run(":8080")
}
