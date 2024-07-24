package repository

import (
	"checkit/models"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func FindEmails() models.SearchResponse {
	query := `{
        "search_type": "matchall",
        "from": 0,
        "max_results": 20,
        "_source": []
    }`
	req, err := http.NewRequest("POST", "http://localhost:4080/api/emails/_search", strings.NewReader(query))
	if err != nil {
		log.Fatal(err)
	}
	req.SetBasicAuth("admin", "Complexpass#123")
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

	var searchResponse models.SearchResponse
	// Unmarshal JSON byte array into Response struct
	err = json.Unmarshal(body, &searchResponse)
	if err != nil {
		log.Fatal(err)
	}

	return searchResponse
}

func FindEmailsBySearch(search string) models.SearchResponse {

	// Define the query structure as a Go map
	query := map[string]interface{}{
		"search_type": "match",
		"query": map[string]interface{}{
			"term": search,
		},
		"from":        0,
		"max_results": 20,
		"_source":     []string{}, // Empty array of strings
	}

	// Convert the query map to JSON
	queryJSON, err := json.Marshal(query)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		log.Fatal(err)
	}

	req, err := http.NewRequest("POST", "http://localhost:4080/api/emails/_search", strings.NewReader(string(queryJSON)))
	if err != nil {
		log.Fatal(err)
	}
	req.SetBasicAuth("admin", "Complexpass#123")
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

	var searchResponse models.SearchResponse
	// Unmarshal JSON byte array into Response struct
	err = json.Unmarshal(body, &searchResponse)
	if err != nil {
		log.Fatal(err)
	}

	return searchResponse
}
