package controllers

import (
	"net/http"

	"github.com/bigzirook/movie-ticket-booking/internal/models"
	"github.com/bigzirook/movie-ticket-booking/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type PaymentController struct {
	DB *gorm.DB
}

// InitiatePayment handles the ticket payment process
func (pc *PaymentController) InitiatePayment(c *gin.Context) {
	var ticket models.Ticket
	ticketID := c.Param("id") // Ticket ID passed as a URL parameter

	// Find the ticket in the database
	if err := pc.DB.Where("id = ?", ticketID).First(&ticket).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Ticket not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch ticket"})
		}
		return
	}

	// Ensure the ticket has not already been paid
	if ticket.Paid {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ticket has already been paid"})
		return
	}

	// Simulate a payment process (use a real payment gateway in a production system)
	paymentResult, err := services.ProcessPayment(ticket.Amount)
	if err != nil {
		c.JSON(http.StatusPaymentRequired, gin.H{"error": "Payment failed", "details": err.Error()})
		return
	}

	// Update ticket status and payment status in the database
	ticket.Status = "booked"
	ticket.Paid = true
	if err := pc.DB.Save(&ticket).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update ticket status"})
		return
	}

	// Respond with payment success
	c.JSON(http.StatusOK, gin.H{
		"message": "Payment successful",
		"payment": paymentResult,
		"ticket":  ticket,
	})
}
