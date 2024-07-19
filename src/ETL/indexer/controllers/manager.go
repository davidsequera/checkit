package controllers

import (
	"fmt"
	"indexer/entities"
	"os"
	"path/filepath"
	"sync"
)

func setParameters() (int, int) {
	// readerWorkes, senderWorkers
	return 30000, 30000
}

func fillPaths(arr *[]string, dir string) error {
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			*arr = append(*arr, path)
		}
		return nil
	})
	return err
}

func Manage(dir string, url string) {
	nRW, nSG := setParameters()

	var wRg sync.WaitGroup
	semaphoreR := make(chan struct{}, nRW)
	var wSg sync.WaitGroup
	semaphoreS := make(chan struct{}, nSG)
	var paths []string
	chEmails := make(chan entities.Email)

	err := fillPaths(&paths, dir)
	if err != nil {
		fmt.Printf("Error walking the directory: %v\n", err)
		return
	}
	fmt.Println("Paths filled")

	go func() {
		for _, path := range paths {
			wRg.Add(1)
			go ReadEmail(path, &wRg, chEmails, semaphoreR)
		}
	}()
	fmt.Println("Reading emails")

	go func() {
		var batch []entities.Email
		for email := range chEmails {
			batch = append(batch, email)
			if len(batch) >= 3000 {
				// "admin", "Complexpass#123"
				emailSender := EmailSender{Url: url, Emails: batch, Credential: struct {
					username string
					password string
				}{username: "admin", password: "Complexpass#123"}}
				batch = []entities.Email{}
				wSg.Add(1)
				go emailSender.SendEmails(&wSg, semaphoreS)
			}
		}
		// "admin", "Complexpass#123"
		emailSender := EmailSender{Url: url, Emails: batch, Credential: struct {
			username string
			password string
		}{username: "admin", password: "Complexpass#123"}}
		wSg.Add(1)
		go emailSender.SendEmails(&wSg, semaphoreS)
	}()

	wRg.Wait()
	fmt.Println("Stop reading emails")

	wSg.Wait()
	fmt.Println("Stop sending emails")

	close(chEmails)
}
