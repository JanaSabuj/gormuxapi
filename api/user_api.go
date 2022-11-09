package api

import (
	"encoding/json"
	"github.com/JanaSabuj/gormuxapi/models"
	"net/http"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var users []models.User
	models.DB.Find(&users)
	json.NewEncoder(w).Encode(users)
}
