package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/0xNicoo/go-microservices/product-service/models"
	"github.com/gorilla/mux"
)

var products = []models.Product{
	{ID: 1, Name: "Laptop", Stock: 10},
	{ID: 2, Name: "Teléfono", Stock: 15},
	{ID: 3, Name: "Teclado", Stock: 20},
	{ID: 4, Name: "Mouse", Stock: 25},
	{ID: 5, Name: "Monitor", Stock: 8},
	{ID: 6, Name: "Impresora", Stock: 5},
	{ID: 7, Name: "Tablet", Stock: 12},
}

func GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func ReduceStock(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID no válido", http.StatusBadRequest)
		return
	}

	var requestBody struct {
		Quantity int `json:"quantity"`
	}

	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	for i, product := range products {
		if product.ID == id {
			if product.Stock < requestBody.Quantity {
				http.Error(w, "Stock insuficiente", http.StatusBadRequest)
			}

			products[i].Stock -= requestBody.Quantity
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(products[i])
			return
		}

	}

	http.Error(w, "No se encontro el producto", http.StatusNotFound)
}
