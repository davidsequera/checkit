package repository

import (
	"checkit/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func FindEmails() (models.SearchResponse, error) {
	config := models.GetConfig()

	fmt.Println("[FindEmails]", config.Username, config.Password)
	query := `{
        "search_type": "matchall",
        "from": 0,
        "max_results": 20,
        "_source": []
    }`
	req, err := http.NewRequest("POST", config.Url+"/_search", strings.NewReader(query))
	if err != nil {
		return models.SearchResponse{}, fmt.Errorf("error creating request: %w", err)
	}
	req.SetBasicAuth(config.Username, config.Password)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return models.SearchResponse{}, fmt.Errorf("error performing request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return models.SearchResponse{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return models.SearchResponse{}, fmt.Errorf("error reading response body: %w", err)
	}

	var searchResponse models.SearchResponse
	err = json.Unmarshal(body, &searchResponse)
	if err != nil {
		return models.SearchResponse{}, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return searchResponse, nil
}

func FindEmailsBySearch(search string) (models.SearchResponse, error) {
	config := models.GetConfig()
	query := map[string]interface{}{
		"search_type": "match",
		"query": map[string]interface{}{
			"term": search,
		},
		"from":        0,
		"max_results": 20,
		"_source":     []string{},
	}

	queryJSON, err := json.Marshal(query)
	if err != nil {
		return models.SearchResponse{}, fmt.Errorf("error marshalling JSON: %w", err)
	}

	req, err := http.NewRequest("POST", config.Url+"/_search", strings.NewReader(string(queryJSON)))
	if err != nil {
		return models.SearchResponse{}, fmt.Errorf("error creating request: %w", err)
	}
	req.SetBasicAuth(config.Username, config.Password)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return models.SearchResponse{}, fmt.Errorf("error performing request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return models.SearchResponse{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return models.SearchResponse{}, fmt.Errorf("error reading response body: %w", err)
	}

	var searchResponse models.SearchResponse
	err = json.Unmarshal(body, &searchResponse)
	if err != nil {
		return models.SearchResponse{}, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return searchResponse, nil
}
