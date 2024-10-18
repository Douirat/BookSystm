package main

import (
	"log"
	"net/http"
	"github.com/Douirat/BookSystm.git/pkg/routes"
	"github.com/gorilla/mux"
	// "gorm.io/driver/mysql"
	// "gorm.io/gorm"
)

func main() {
router := mux.NewRouter()
routes.BookSystemRouter(router)
http.Handle("/", router)
log.Fatal(http.ListenAndServe(":8080", router))
}