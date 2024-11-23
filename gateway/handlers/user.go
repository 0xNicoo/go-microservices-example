package handlers

import (
	"io"
	"net/http"
)

func UserHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://user-service:8081/users")
	if err != nil {
		http.Error(w, "Error comunicándose con el servicio de usuarios", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}
