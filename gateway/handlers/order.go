package handlers

import (
	"io"
	"net/http"
)

func CreateOrderHandler(w http.ResponseWriter, r *http.Request) {
	url := "http://order-service:8083/orders"

	//Crea el request
	req, err := http.NewRequest("POST", url, r.Body)
	if err != nil {
		http.Error(w, "Error creando la solicitud", http.StatusInternalServerError)
		return
	}
	req.Header = r.Header

	// Hace la solicitud
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "Error al contactar el microservicio", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}

func GetOrdersHandler(w http.ResponseWriter, r *http.Request) {
	url := "http://order-service:8083/orders"

	//Crea el request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		http.Error(w, "Error creando la solicitud", http.StatusInternalServerError)
		return
	}
	req.Header = r.Header

	// Hace la solicitud
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "Error al contactar el microservicio", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}
