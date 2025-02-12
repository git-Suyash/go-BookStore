package routes

import (
	"github.com/git-Suyash/go-BookStore/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/api/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/api/book/", controllers.GetBook).Methods("GET")
	router.HandleFunc("/api/book/{id}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/api/book/{id}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/api/book/{id}", controllers.DeleteBook).Methods("DELETE")
}
