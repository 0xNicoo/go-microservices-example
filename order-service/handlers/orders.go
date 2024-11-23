package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/0xNicoo/go-microservices/order-service/models"
)

var (
	orders     []models.Order
	orderMutex sync.Mutex
	nextID     = 1
)

const (
	UserServiceURL    = "http://user-service:8081/users"
	ProductServiceURL = "http://product-service:8082"
)

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	var newOrder models.Order

	if err := json.NewDecoder(r.Body).Decode(&newOrder); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	reduceStockURL := fmt.Sprintf("%s/products/%d/reduce-stock", ProductServiceURL, newOrder.ProductID)

	requestBody := struct {
		Quantity int `json:"quantity"`
	}{
		Quantity: newOrder.Quantity,
	}

	requestBodyJSON, err := json.Marshal(requestBody)
	if err != nil {
		http.Error(w, "Error preparando la solicitud al servicio de productos", http.StatusInternalServerError)
		return
	}

	req, err := http.NewRequest("PUT", reduceStockURL, bytes.NewBuffer(requestBodyJSON))
	if err != nil {
		http.Error(w, "Error creando la solicitud al servicio de productos", http.StatusInternalServerError)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "Error al contactar el servicio de productos", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		http.Error(w, "Error del servicio de productos: "+resp.Status, resp.StatusCode)
		return
	}

	newOrder.ID = nextID
	newOrder.CreatedAt = time.Now()

	orderMutex.Lock()
	orders = append(orders, newOrder)
	nextID++
	orderMutex.Unlock()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newOrder)
}

func GetOrders(w http.ResponseWriter, r *http.Request) {
	orderMutex.Lock()
	defer orderMutex.Unlock()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(orders)
}
