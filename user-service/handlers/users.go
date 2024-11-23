package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/0xNicoo/go-microservices/user-service/models"
	"github.com/gorilla/mux"
)

var users = []models.User{
	{ID: 1, Name: "Nicolas"},
	{ID: 2, Name: "Bautista"},
	{ID: 3, Name: "Matias"},
	{ID: 4, Name: "Federico"},
	{ID: 5, Name: "Tomas"},
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID no válido", http.StatusBadRequest)
		return
	}

	for i, user := range users {
		if user.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(users[i])
			return
		}
	}

	http.Error(w, "No se encontro el usuario", http.StatusNotFound)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
