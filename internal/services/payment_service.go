package services

import (
	"fmt"
	"math/rand"
	"time"
)

// Dummy payment service
func ProcessPayment(amount int) (string, error) {
	rand.Seed(time.Now().UnixNano())
	if rand.Intn(2) == 0 {
		return "", fmt.Errorf("Payment failed")
	}
	return "Payment successful", nil
}
