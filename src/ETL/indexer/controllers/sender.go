package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"indexer/entities"
	"io"
	"log"
	"net/http"
	"sync"
)

type EmailSender struct {
	Url        string
	Emails     []entities.Email
	Credential struct {
		username string
		password string
	}
}

func (es *EmailSender) SendEmails(wg *sync.WaitGroup, semaphore chan struct{}) {

	for _, email := range es.Emails {
		fmt.Println(email.MessageID)
	}
	// Outer structure
	payload := struct {
		Index   string           `json:"index"`
		Records []entities.Email `json:"records"`
	}{
		Index:   "devemails",
		Records: es.Emails,
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest("POST", es.Url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		log.Fatal(err)
	}
	// "admin", "Complexpass#123"
	req.SetBasicAuth(es.Credential.username, es.Credential.password)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	log.Println(resp.StatusCode)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))

	defer wg.Done()

}
