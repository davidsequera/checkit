package controller

import (
	"checkit/services"
	"net/http"
)

func CreateEmail(w http.ResponseWriter, r *http.Request) {
	services.CreateEmail(w, r)
}

func GetEmails(w http.ResponseWriter, r *http.Request) {
	services.GetEmails(w, r)
}

func GetEmail(w http.ResponseWriter, r *http.Request) {
	services.GetEmail(w, r)
}

func UpdateEmail(w http.ResponseWriter, r *http.Request) {
	services.UpdateEmail(w, r)
}

func DeleteEmail(w http.ResponseWriter, r *http.Request) {
	services.DeleteEmail(w, r)
}
