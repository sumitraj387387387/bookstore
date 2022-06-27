package controllers

import (
	"bookProject/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

var books []models.Book

func ErrorHandler(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
func AddBooks() {
	books = append(books, models.Book{Id: "1", Title: "Differential Calculus", Author: &models.Author{FirstName: "Sumit", LastName: "Raj"}})
	books = append(books, models.Book{Id: "2", Title: "Integral Calculus", Author: &models.Author{FirstName: "Sameer", LastName: "Raj"}})
	books = append(books, models.Book{Id: "3", Title: "Complex Number", Author: &models.Author{FirstName: "Ankit", LastName: "Agarwal"}})
	books = append(books, models.Book{Id: "4", Title: "World Affairs", Author: &models.Author{FirstName: "Prashant", LastName: "Dhawan"}})
}

func GetBooks(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-type", "application/json")
	err := json.NewEncoder(w).Encode(books)
	ErrorHandler(err)

}

func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	for _, item := range books {
		if item.Id == params["id"] {
			err := json.NewEncoder(w).Encode(item)
			ErrorHandler(err)
			return
		}
	}
	err := json.NewEncoder(w).Encode(&models.Book{})
	ErrorHandler(err)

}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var book models.Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.Id = strconv.Itoa(rand.Intn(10000))
	books = append(books, book)
	err := json.NewEncoder(w).Encode(book)
	ErrorHandler(err)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.Id == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book models.Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.Id = params["id"]
			books = append(books, book)
			err := json.NewEncoder(w).Encode(book)
			ErrorHandler(err)
			return
		}
	}
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.Id == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	err := json.NewEncoder(w).Encode(books)
	ErrorHandler(err)
}
