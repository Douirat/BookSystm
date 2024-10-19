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
func GetAllBooks() []Book {
	var Books []Book
	connection.Find(&Books)
	return Books
}

// Get a specific book by its id:
func GetBookById(id int64) (*Book, *gorm.DB) {
	var book Book
	connection = connection.Where("ID=?", id).Find(&book)
	return &book, connection
}

// Create a new instance of a book, returns true if created successfully, false otherwise.
func CreateBook(book *Book) bool {
	connection.Create(book) // Call the Create method
	// Check if the book's ID was set, indicating a successful creation
	return book.ID != 0
}

// Update info withing an existing instance of a book:
func UpdateBook(id int64, Auth string) {
	var book Book
	connection.Model(&book).Where("ID=?", id).Update("Author", Auth)
}

// Delete a book with a specific id:
func DeletBook(id int64) Book {
	var book Book
	connection.Where("ID=?", id).Delete(book)
	return book
}
