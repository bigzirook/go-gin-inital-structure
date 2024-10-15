package routes

import (
	"github.com/bigzirook/movie-ticket-booking/config"
	"github.com/bigzirook/movie-ticket-booking/internal/controllers"
	"github.com/bigzirook/movie-ticket-booking/internal/middlewares"
	_ "github.com/bigzirook/movie-ticket-booking/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	// Initialize the database
	config.ConnectDB()

	authController := controllers.AuthController{DB: config.DB}
	bookingController := controllers.BookingController{DB: config.DB}
	eventController := controllers.EventController{DB: config.DB}
	// Public routes
	router.POST("/register", authController.Register)
	router.POST("/login", authController.Login)

	// Secured routes with JWT middleware
	protected := router.Group("/api")
	protected.Use(middlewares.AuthMiddleware())
	{
		// Event routes
		protected.GET("/events", eventController.ListEvents)   // List all events
		protected.GET("/events/:id", eventController.GetEvent) // Get event details by ID

		// Booking routes
		protected.POST("/book-ticket", bookingController.BookTicket)
	}
}
