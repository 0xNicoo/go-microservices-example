package handlers

import (
	"io"
	"net/http"
)

func ProductHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://product-service:8082/products")
	if err != nil {
		http.Error(w, "Error comunicándose con el servicio de productos", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}
