package models

import "github.com/jinzhu/gorm"

// Ticket represents a ticket booked by a user for an event
type Ticket struct {
	gorm.Model
	UserID  uint   `json:"user_id"`
	EventID uint   `json:"event_id"`
	Status  string `json:"status"` // Booking status (e.g., booked, canceled)
	Amount  int    `json:"amount"` // Payment amount
	Paid    bool   `json:"paid"`   // Payment status (paid or not)
}
