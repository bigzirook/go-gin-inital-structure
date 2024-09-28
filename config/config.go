package config

import (
    "fmt"
    "log"
    "os"

    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "github.com/joho/godotenv"
)

var DB *gorm.DB

// LoadEnv loads environment variables from .env file
func LoadEnv() {
    err := godotenv.Load()
    if err != nil {
        log.Println("No .env file found, using default environment variables")
    }
}

// ConnectDB initializes the database connection using environment variables
func ConnectDB() {
    LoadEnv()

    // Get environment variables or set default values
    dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")
    dbName := os.Getenv("DB_NAME")

    // Format the database connection string
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", 
        dbUser, dbPassword, dbHost, dbPort, dbName)

    var err error
    DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
       
