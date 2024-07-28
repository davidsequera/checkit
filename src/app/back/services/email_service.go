package services

import (
	"checkit/models"
	"checkit/repository"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

var mock_emails = []models.Email{
	{MessageID: "1", From: "olivia@example.com", Subject: "Meeting Agenda", Date: "2023-06-01"},
	{MessageID: "2", From: "john@example.com", Subject: "Quarterly Report", Date: "2023-05-15"},
	{MessageID: "3", From: "emily@example.com", Subject: "New Project Update", Date: "2023-04-30"},
	{MessageID: "4", From: "michael@example.com", Subject: "Feedback Request", Date: "2023-04-20"},
	{MessageID: "5", From: "sophia@example.com", Subject: "Invoice Reminder", Date: "2023-03-28"},
	{MessageID: "6", From: "david@example.com", Subject: "Team Announcement", Date: "2023-03-15"},
	{MessageID: "7", From: "isabella@example.com", Subject: "Upcoming Event", Date: "2023-02-28"},
	{MessageID: "8", From: "alexander@example.com", Subject: "Budget Approval", Date: "2023-02-10"},
	{MessageID: "9", From: "ava@example.com", Subject: "New Hire Announcement", Date: "2023-01-25"},
	{MessageID: "10", From: "william@example.com", Subject: "Expense Report", Date: "2023-01-05"},
}

func GetEmails(w http.ResponseWriter, r *http.Request) {
	var searchResponse models.SearchResponse
	var err error
	query := r.URL.Query()
	search := query.Get("search")

	if search != "" {
		searchResponse, err = repository.FindEmailsBySearch(search)
	} else {
		searchResponse, err = repository.FindEmails()
	}

	if err != nil {
		log.Println(err)
		w.Write([]byte{'E', 'r', 'r', 'o', 'r'})
	}

	var emails []models.Email = []models.Email{}
	for _, hint := range searchResponse.Hits.Hits {
		emails = append(emails, hint.Source)
	}

	json.NewEncoder(w).Encode(emails)
}

func GetEmail(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	for _, email := range mock_emails {
		if email.MessageID == id {
			json.NewEncoder(w).Encode(email)
			return
		}
	}
	http.NotFound(w, r)
}

func CreateEmail(w http.ResponseWriter, r *http.Request) {
	var email models.Email
	json.NewDecoder(r.Body).Decode(&email)
	email.MessageID = fmt.Sprintf("%d", len(mock_emails)+1)
	mock_emails = append(mock_emails, email)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(email)
}

func UpdateEmail(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var updatedEmail models.Email
	json.NewDecoder(r.Body).Decode(&updatedEmail)

	for i, email := range mock_emails {
		if email.MessageID == id {
			updatedEmail.MessageID = id
			mock_emails[i] = updatedEmail
			json.NewEncoder(w).Encode(updatedEmail)
			return
		}
	}
	http.NotFound(w, r)
}

func DeleteEmail(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	for i, email := range mock_emails {
		if email.MessageID == id {
			mock_emails = append(mock_emails[:i], mock_emails[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.NotFound(w, r)
}
