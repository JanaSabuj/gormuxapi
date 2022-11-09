package main

import (
	"fmt"
	"github.com/gorilla/mux"
)

func InitRouter() {
	r := mux.NewRouter()
	r.HandleFunc("/users", GetUsers).Methods("GET")
	r.HandleFunc("/users/{id}", GetUser).Methods("GET")
	r.HandleFunc("/users", CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", DeleteUser).Methods("DELETE")
	fmt.Println(r)
}

func main() {
	fmt.Println("Hello to gormuxapi")
	InitRouter()
}
