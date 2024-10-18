package models

import (
	"log"

	"github.com/Douirat/BookSystm.git/pkg/configs"
	"gorm.io/gorm"
)

// Declare a book object:
type Book struct {
	gorm.Model
	Title       string `json:"title"`
	Author      string `json:"author"`
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

// Get all the books from database:
func GetAllBook() {
}

// Get a specific book by its id:
func GetBookById() {
}

// Create a new instance of a book, returns true if created successfully, false otherwise.
func CreateBook(book *Book) bool {
	connection.Create(book) // Call the Create method
	// Check if the book's ID was set, indicating a successful creation
	return book.ID != 0
}

// Update info withing an existing instance of a book:
func UpdateBook() {
}

// Delete a book with a specific id:
func DeletBook() {
}
