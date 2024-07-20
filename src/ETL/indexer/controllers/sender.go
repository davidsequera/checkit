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

func SendEmails(url string, emails []entities.Email, user *string, password *string, wg *sync.WaitGroup, semaphore chan struct{}) {
	semaphore <- struct{}{}        // Acquire semaphore
	defer func() { <-semaphore }() // Release semaphore
	defer wg.Done()
	// Outer structure
	payload := struct {
		Index   string           `json:"index"`
		Records []entities.Email `json:"records"`
	}{
		Index:   "emails",
		Records: emails,
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		log.Fatal(err)
	}
	req.SetBasicAuth(*user, *password)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp.StatusCode)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))

	resp.Body.Close()
}
