package main

import (
	"log"
	"net/http"

	"github.com/0xNicoo/go-microservices/product-service/handlers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/products", handlers.GetProducts).Methods("GET")
	r.HandleFunc("/products/{id}/reduce-stock", handlers.ReduceStock).Methods("PUT")

	log.Println("Product Service corriendo en el puerto 8082...")
	log.Fatal(http.ListenAndServe(":8082", r))
}
