package controllers

import (
	"net/http"

	"github.com/bigzirook/movie-ticket-booking/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type BookingController struct {
	DB *gorm.DB
}

// List available events
func (bc *BookingController) ListEvents(c *gin.Context) {
	var events []models.Event
	bc.DB.Find(&events)
	c.JSON(http.StatusOK, events)
}

// Book a ticket
func (bc *BookingController) BookTicket(c *gin.Context) {
	var ticket models.Ticket
	if err := c.ShouldBindJSON(&ticket); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Save the ticket in DB
	bc.DB.Create(&ticket)
	c.JSON(http.StatusOK, gin.H{"message": "Ticket booked successfully!"})
}
