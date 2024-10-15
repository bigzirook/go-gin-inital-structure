package models

import "github.com/jinzhu/gorm"

type Event struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       int    `json:"price"`
}
