package main

import (
	"fmt"
	"github.com/JanaSabuj/gormuxapi/api"
	"github.com/JanaSabuj/gormuxapi/models"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func InitRouter() {
	r := mux.NewRouter()

	UserApiInstance := api.NewUserApi()
	r.HandleFunc("/users", UserApiInstance.GetUsers).Methods("GET")
	r.HandleFunc("/users/{id}", UserApiInstance.GetUser).Methods("GET")
	r.HandleFunc("/users", UserApiInstance.CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", UserApiInstance.UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", UserApiInstance.DeleteUser).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":9000", r))
}

func main() {
	fmt.Println("Hello to gormuxapi")
	models.InitDBMigration()
	InitRouter()
}
