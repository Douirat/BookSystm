package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Douirat/BookSystm.git/pkg/models"
	"github.com/Douirat/BookSystm.git/pkg/utils"
	"github.com/gorilla/mux"
)

// Declare my book struct:
var NewBook *models.Book

// Get all the books from database:
func GetAllBook(wr http.ResponseWriter, rq *http.Request) {
	books := models.GetAllBooks()
	response, _ := json.Marshal(books)
	wr.Header().Set("Content-Type", "Application/json")
	wr.WriteHeader(http.StatusOK)
	wr.Write(response)
}

// Get a specific book by its id:
func GetBookById(wr http.ResponseWriter, rq *http.Request) {
}

// Create a new instance of a book:
func CreateBook(wr http.ResponseWriter, rq *http.Request) {
}

// Update info withing an existing instance of a book:
func UpdateBook(wr http.ResponseWriter, rq *http.Request) {
}

// Delete a book with a specific id:
func DeletBook(wr http.ResponseWriter, rq *http.Request) {
}
