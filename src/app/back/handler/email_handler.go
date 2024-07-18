package handlers

import (
	"checkit/services"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type EmailHandler struct {
	service *services.EmailService
}

func NewEmailHandler(service *services.EmailService) *EmailHandler {
	return &EmailHandler{service: service}
}

func (h *EmailHandler) GetEmails(w http.ResponseWriter, r *http.Request) {
	emails, err := h.service.GetEmails()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(emails)
}

func (h *EmailHandler) GetEmail(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	email, err := h.service.GetEmail(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if email == nil {
		http.NotFound(w, r)
		return
	}
	json.NewEncoder(w).Encode(email)
}
