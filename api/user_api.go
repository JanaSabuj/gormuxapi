package api

import (
	"encoding/json"
	"fmt"
	"github.com/JanaSabuj/gormuxapi/models"
	"github.com/gorilla/mux"
	"net/http"
)

type UserApi struct {
}

func NewUserApi() UserApiInterface {
	return &UserApi{}
}

func (self *UserApi) GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var users []models.User
	models.DB.Find(&users)
	json.NewEncoder(w).Encode(users)
}

func (self *UserApi) GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.User
	params := mux.Vars(r) // route params
	models.DB.First(&user, params["id"])
	json.NewEncoder(w).Encode(user)
}

func (self *UserApi) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.User
	json.NewDecoder(r.Body).Decode(&user) // get it
	fmt.Println(user)
	fmt.Println(models.DB)
	models.DB.Create(&user)         // save it
	json.NewEncoder(w).Encode(user) // return it
}

func (self *UserApi) UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var user models.User
	models.DB.First(&user, params["id"])
	json.NewDecoder(r.Body).Decode(&user)
	models.DB.Save(&user)
	json.NewEncoder(w).Encode(user)
}

func (self *UserApi) DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var user models.User
	models.DB.Delete(&user, params["id"])
	json.NewEncoder(w).Encode("The USer is Deleted Successfully!")
}
