package main

import (
	"github.com/bigzirook/movie-ticket-booking/config"

	"github.com/bigzirook/movie-ticket-booking/internal/models"
	"github.com/bigzirook/movie-ticket-booking/internal/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize environment
	config.LoadEnv()

	// Initialize database connection
	config.ConnectDB()

	// Auto-migrate database models
	config.DB.AutoMigrate(&models.User{}, &models.Event{}, &models.Ticket{})
	// config.DB.AutoMigrate(&models.User{})

	// Initialize the Gin router
	r := gin.Default()
	// fmt.Print(r)

	// Setup routes
	routes.SetupRoutes(r)

	// Run the server
	r.Run(":8080")
}
