package main

import (
	"log"
	"net/http"

	"github.com/0xNicoo/go-microservices/user-service/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/users", handlers.GetUsers).Methods("GET")

	log.Println("User Service corriendo en el puerto 8081...")
	log.Fatal(http.ListenAndServe(":8081", r))
}
