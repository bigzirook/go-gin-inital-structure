package config

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql" // MySQL driver
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

var DB *gorm.DB

// LoadEnv loads environment variables
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
}

// ConnectDB connects to the MySQL database
func ConnectDB() {
	var err error

	// MySQL connection string: "user:password@tcp(127.0.0.1:3306)/dbname?charset=utf8&parseTime=True&loc=Local"
	connStr := os.Getenv("DB_URL")

	DB, err = gorm.Open("mysql", connStr)
	if err != nil {
		panic("failed to connect to database: " + err.Error())
	}
	fmt.Println("Database connected!")
}
