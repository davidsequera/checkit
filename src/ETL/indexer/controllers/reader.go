package controllers

import (
	"fmt"
	"indexer/entities"
	"os"
	"sync"
)

func ReadEmail(path string, wg *sync.WaitGroup, ch chan<- entities.Email, semaphore chan struct{}) {
	semaphore <- struct{}{}        // Acquire semaphore
	defer func() { <-semaphore }() // Release semaphore
	defer wg.Done()
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Cagamos HPTASSSSS", path)
		// fmt.Printf("Error reading file %s: %v\n", path, err)
		return
	}
	email := entities.ParseEmail(string(data))
	ch <- email
}
