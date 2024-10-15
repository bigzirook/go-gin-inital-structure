package services

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/bigzirook/movie-ticket-booking/internal/models"

	"github.com/dgrijalva/jwt-go"
)

// JWTClaims struct contains the user details and standard JWT claims
type JWTClaims struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	UserID   uint   `json:"user_id"`
	jwt.StandardClaims
}

// SecretKey retrieves the JWT secret from the environment
func SecretKey() string {
	return os.Getenv("JWT_SECRET")
}

// GenerateJWT generates a JWT token for the authenticated user
func GenerateJWT(user models.User) (string, error) {
	// Define the expiration time for the token
	expirationTime := time.Now().Add(24 * time.Hour) // Token valid for 24 hours

	// Create the JWT claims, including the user information
	claims := &JWTClaims{
		Username: user.Username,
		Email:    user.Email,
		UserID:   user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(), // Set expiration time in the claims
		},
	}

	// Create the token using the claims and sign it with the secret key
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(SecretKey()))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

// ValidateJWT parses and validates the JWT token
func ValidateJWT(tokenString string) (*JWTClaims, error) {
	// Parse the token and check if it is valid
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Check if the signing method is HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(SecretKey()), nil
	})

	// If parsing or validation failed, return the error
	if err != nil {
		return nil, err
	}

	// Extract the claims and check if the token is valid
	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
