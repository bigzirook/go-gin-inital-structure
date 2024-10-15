package controllers

import (
	"net/http"

	"github.com/bigzirook/movie-ticket-booking/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type EventController struct {
	DB *gorm.DB
}

// ListEvents returns all available events
func (ec *EventController) ListEvents(c *gin.Context) {
	var events []models.Event
	if err := ec.DB.Find(&events).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch events"})
		return
	}
	c.JSON(http.StatusOK, events)
}

// CreateEvent adds a new event (admin access)
func (ec *EventController) CreateEvent(c *gin.Context) {
	var event models.Event
	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Save the new event in the database
	if err := ec.DB.Create(&event).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create event"})
		return
	}

	c.JSON(http.StatusCreated, event)
}

// GetEvent retrieves the details of a specific event by ID
func (ec *EventController) GetEvent(c *gin.Context) {
	var event models.Event
	id := c.Param("id")

	// Fetch the event by ID
	if err := ec.DB.Where("id = ?", id).First(&event).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch event"})
		}
		return
	}

	c.JSON(http.StatusOK, event)
}

// UpdateEvent updates an existing event (admin access)
func (ec *EventController) UpdateEvent(c *gin.Context) {
	var event models.Event
	id := c.Param("id")

	// Fetch the event by ID
	if err := ec.DB.Where("id = ?", id).First(&event).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch event"})
		}
		return
	}

	// Bind the JSON data to the event
	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update the event in the database
	if err := ec.DB.Save(&event).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update event"})
		return
	}

	c.JSON(http.StatusOK, event)
}

// DeleteEvent removes an event by ID (admin access)
func (ec *EventController) DeleteEvent(c *gin.Context) {
	id := c.Param("id")

	// Delete the event from the database
	if err := ec.DB.Where("id = ?", id).Delete(&models.Event{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete event"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully"})
}
