package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"time"
)

func searchAndCount(pattern *regexp.Regexp, dir string) (map[string]int, error) {
	matches := make(map[string]int) // Map to store pattern:count

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil // Skip directories
		}

		// Read file content
		data, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("Error reading file %s: %w", path, err)
		}

		allMatches := pattern.FindAllString(string(data), -1)
		for _, match := range allMatches {
			matches[match] = matches[match] + 1 // Increment count for each match
		}

		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("Error walking directory: %w", err)
	}

	return matches, nil
}

func saveMapToFile(data map[string]int, filename string) error {
	// Marshal the map to JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("Error marshalling map to JSON: %w", err)
	}

	// Create and write to file
	err = os.WriteFile(filename, jsonData, 0644) // Set appropriate permissions
	if err != nil {
		return fmt.Errorf("Error writing map to file: %w", err)
	}

	return nil
}

func main() {
	startTime := time.Now()

	patternString := `(?m)^([^:\s]*):` // Capture everything before colon

	// Compile the pattern
	pattern, err := regexp.Compile(patternString)
	if err != nil {
		fmt.Println("Error compiling pattern:", err)
		return
	}

	// Specify directory to search
	searchDir := "db"

	// Search and count matches
	matchCounts, err := searchAndCount(pattern, searchDir)
	if err != nil {
		fmt.Println("Error searching files:", err)
		return
	}

	// Save map to file (replace with desired filename)
	err = saveMapToFile(matchCounts, "src/explorer/match_counts.json")
	if err != nil {
		fmt.Println("Error saving map to file:", err)
		return
	}

	fmt.Println("Match counts saved to match_counts.json")

	endTime := time.Now()
	elapsedTime := endTime.Sub(startTime)
	fmt.Println("Elapsed time:", elapsedTime)
}
