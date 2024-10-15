package controllers

import (
	"fmt"
	"net/http"

	"github.com/bigzirook/movie-ticket-booking/config"
	"github.com/bigzirook/movie-ticket-booking/internal/models"
	"github.com/bigzirook/movie-ticket-booking/internal/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/gorm"
)

type AuthController struct {
	DB *gorm.DB
}

var validate = validator.New()

// Register a new user
func (ac *AuthController) Register(c *gin.Context) {
	var user models.User
	// fmt.Println("Attempting to bind incoming JSON to the User model")
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate the user struct fields
	if err := validate.Struct(user); err != nil {
		// Handle validation errors
		validationErrors := err.(validator.ValidationErrors)
		errors := make(map[string]string)
		for _, fieldError := range validationErrors {
			errors[fieldError.Field()] = fieldError.Error() // Return default error message
		}

		c.JSON(http.StatusBadRequest, gin.H{"validation_errors": errors})
		return
	}

	// Check if user with the same email exists
	var existingUser models.User
	if err := config.DB.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already in use"})
		return
	}

	// Hash the user's password before storing it
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	user.Password = hashedPassword // Store the hashed password

	// Save the new user to the database
	if err := config.DB.Create(&user).Error; err != nil {
		// Handle database errors (e.g., unique constraint violation)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	fmt.Println("Deepak====>", ac.DB.Create(&user).Error)
	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully!"})
}

// User login with JWT
/*func (ac *AuthController) Login(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check user credentials (simplified, you should hash passwords)
	if err := ac.DB.Where("username = ? AND password = ?", user.Username, user.Password).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Generate JWT token
	token, err := services.GenerateJWT(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Token generation failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
} */

// Login user and generate JWT token
func (ac *AuthController) Login(c *gin.Context) {
	var loginData struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Find the user by email
	var user models.User

	fmt.Println(user.Password, loginData.Password)
	if err := config.DB.Where("email = ?", loginData.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Check if the password matches
	if !utils.CheckPasswordHash(user.Password, loginData.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Generate JWT token using the utility function
	token, err := utils.GenerateJWT(user.ID, user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// Return JWT token to the client
	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful!",
		"token":   token,
	})
}
