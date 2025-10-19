package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"example.com/go-api-crud/models"
	"example.com/go-api-crud/routes"
	"example.com/go-api-crud/config"
)


func jsonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func main() {
	config.ConnectDB()
	config.DB.AutoMigrate(&models.User{})

	r := mux.NewRouter()

	r.Use(jsonMiddleware)

	routes.RegisterUserRoutes(r)

	log.Println("Aplikasi berjalan di http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
