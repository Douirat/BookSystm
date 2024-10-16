package models

import (
	"log"
	"github.com/Douirat/BookSystm.git/pkg/configs"
	"gorm.io/gorm"
)

// Declare a book object:
type Book struct {
	gorm.Model
	Title string `json:"title"`
	Author string `json:"author"`
	Publication string `json:"publication"`
}

// Declare a global variable to hold the connection with database:
var connection *gorm.DB

// Initialize a database connection:
func Init() {
	var err error
	connection, err = configs.DataBaseConncetion()
	if err != nil {
		log.Fatal(err)
		return
	}
	connection.AutoMigrate()
}