package models

type SearchResponse struct {
	Took     int     `json:"took"`
	TimedOut bool    `json:"timed_out"`
	MaxScore float64 `json:"max_score"`
	Hits     struct {
		Total struct {
			Value int `json:"value"`
		} `json:"total"`
		Hits []struct {
			Index     string  `json:"_index"`
			Type      string  `json:"_type"`
			ID        string  `json:"_id"`
			Score     float64 `json:"_score"`
			Timestamp string  `json:"@timestamp"`
			Source    Email   `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`
	Buckets interface{} `json:"buckets"`
	Error   string      `json:"error"`
}
