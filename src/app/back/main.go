package main

import (
	"checkit/services"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func main() {

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
	r.Post("/emails", services.CreateEmail)
	// Read
	r.Get("/emails", services.GetEmails)
	r.Get("/emails/{id}", services.GetEmail)
	// Update
	r.Put("/emails/{id}", services.UpdateEmail)
	// Delete
	r.Delete("/emails/{id}", services.DeleteEmail)

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", r)
}
