package routes

import (
	"github.com/gorilla/mux"
	"example.com/go-api-crud/controllers"
)

func RegisterUserRoutes(router *mux.Router) {
	router.HandleFunc("/api/users", controllers.GetUsers).Methods("GET")
	router.HandleFunc("/api/users/{id}", controllers.GetUser).Methods("GET")
	router.HandleFunc("/api/users", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/api/users/{id}", controllers.UpdateUser).Methods("PUT")
	router.HandleFunc("/api/users/{id}", controllers.DeleteUser).Methods("DELETE")
}