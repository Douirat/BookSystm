package routes

import (
	"github.com/Douirat/BookSystm.git/pkg/controllers"	
	"github.com/gorilla/mux"
)

// Create the router:
var BookSystemRouter = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.GetAllBook).Methods("GET")
	router.HandleFunc("/book/{id}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/", controllers.Create).Methods("POST")
	router.HandleFunc("/book/{id}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{id}", controllers.DeletBook).Methods("DELETE")
}