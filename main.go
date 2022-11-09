package main

import (
	"fmt"
	"github.com/JanaSabuj/gormuxapi/api"
	"github.com/gorilla/mux"
)

func InitRouter() {
	r := mux.NewRouter()
	r.HandleFunc("/users", api.GetUsers).Methods("GET")
	//r.HandleFunc("/users/{id}", api.GetUser).Methods("GET")
	//r.HandleFunc("/users", api.CreateUser).Methods("POST")
	//r.HandleFunc("/users/{id}", api.UpdateUser).Methods("PUT")
	//r.HandleFunc("/users/{id}", api.DeleteUser).Methods("DELETE")
	fmt.Println(r)
}

func main() {
	fmt.Println("Hello to gormuxapi")
	InitRouter()
}
