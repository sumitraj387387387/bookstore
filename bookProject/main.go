package main

import (
	"bookProject/controllers"
	"bookProject/routes"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	controllers.AddBooks()
	routes.Setup(r)
	log.Fatal(http.ListenAndServe(":8000", r))

}
