package main

import (
	"log"
	"net/http"

	"github.com/0xNicoo/go-microservices/order-service/handlers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/orders", handlers.CreateOrder).Methods("POST")
	router.HandleFunc("/orders", handlers.GetOrders).Methods("GET")

	log.Println("Order Service running on port 8083")
	http.ListenAndServe(":8083", router)
}
