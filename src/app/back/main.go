package main

import (
	"checkit/controller"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {

	// Enviroment
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}
	r := chi.NewRouter()

	// Configure CORS options
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"}, // Allow all origins
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any major browsers
	}))

	// Middleware
	r.Use(middleware.Logger)

	// Routes
	// CRUD

	// Create
	r.Post("/emails", controller.CreateEmail)
	// Read
	r.Get("/emails", controller.GetEmails)
	r.Get("/emails/{id}", controller.GetEmail)
	// Update
	r.Put("/emails/{id}", controller.UpdateEmail)
	// Delete
	r.Delete("/emails/{id}", controller.DeleteEmail)

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", r)
}
