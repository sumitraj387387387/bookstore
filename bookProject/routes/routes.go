package routes

import (
	"bookProject/controllers"
	"github.com/gorilla/mux"
)

func Setup(r *mux.Router) {

	r.HandleFunc("/books", controllers.GetBooks).Methods("GET")
	r.HandleFunc("/books/{id}", controllers.GetBook).Methods("GET")
	r.HandleFunc("/books", controllers.CreateBook).Methods("POST")
	r.HandleFunc("/books/{id}", controllers.UpdateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", controllers.DeleteBook).Methods("DELETE")
}
