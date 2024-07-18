package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Email struct {
	ID      string `json:"id"`
	Address string `json:"address"`
	Name    string `json:"name"`
}

var emails = []Email{
	{ID: "1", Address: "john@company.com", Name: "John Doe"},
	{ID: "2", Address: "jane@company.com", Name: "Jane Smith"},
}

func main() {
	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.Logger)

	// Routes
	r.Get("/emails", getEmails)
	r.Get("/emails/{id}", getEmail)
	r.Post("/emails", createEmail)
	r.Put("/emails/{id}", updateEmail)
	r.Delete("/emails/{id}", deleteEmail)

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", r)
}

func getEmails(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(emails)
}

func getEmail(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	for _, email := range emails {
		if email.ID == id {
			json.NewEncoder(w).Encode(email)
			return
		}
	}
	http.NotFound(w, r)
}

func createEmail(w http.ResponseWriter, r *http.Request) {
	var email Email
	json.NewDecoder(r.Body).Decode(&email)
	email.ID = fmt.Sprintf("%d", len(emails)+1)
	emails = append(emails, email)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(email)
}

func updateEmail(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var updatedEmail Email
	json.NewDecoder(r.Body).Decode(&updatedEmail)

	for i, email := range emails {
		if email.ID == id {
			updatedEmail.ID = id
			emails[i] = updatedEmail
			json.NewEncoder(w).Encode(updatedEmail)
			return
		}
	}
	http.NotFound(w, r)
}

func deleteEmail(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	for i, email := range emails {
		if email.ID == id {
			emails = append(emails[:i], emails[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.NotFound(w, r)
}
