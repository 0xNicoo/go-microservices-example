package main

import (
	"log"
	"net/http"

	"github.com/0xNicoo/go-microservices/gateway/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Definir rutas
	r.HandleFunc("/api/users", handlers.UserHandler).Methods("GET", "POST")
	r.HandleFunc("/api/products", handlers.ProductHandler).Methods("GET")
	r.HandleFunc("/api/orders", handlers.CreateOrderHandler).Methods("POST")
	r.HandleFunc("/api/orders", handlers.GetOrdersHandler).Methods("GET")

	log.Println("API Gateway corriendo en el puerto 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
