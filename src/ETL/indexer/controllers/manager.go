package controllers

import (
	"fmt"
	"indexer/entities"
	"os"
	"path/filepath"
	"sync"
)

func setParameters(size int) (int, int, int) {
	// readerWorkes, senderWorkers, batchSize
	if size > 20000 {
		return 30000, 300, 1000
	}
	if size <= 2000 {
		return size, size, size
	}
	return size, size / 3, size / 5
}

func fillPaths(dir string) ([]string, error) {
	var arr = []string{}
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			arr = append(arr, path)
		}
		return nil
	})
	return arr, err
}

func Manage(dir string, url string, user *string, password *string) {
	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Printf("Error Reading the root: %v\n", err)
		return
	}
	for _, f := range files {
		if f.IsDir() {
			manageUser(dir+"/"+f.Name(), url, user, password)
		}
	}
}

func manageUser(dir string, url string, user *string, password *string) {
	paths, err := fillPaths(dir)
	if err != nil {
		fmt.Printf("Error Reading the directory: %v\n", err)
		return
	}

	nRW, nSG, batchSize := setParameters(len(paths))
	var wRg sync.WaitGroup
	var wSg sync.WaitGroup
	semaphoreR := make(chan struct{}, nRW)
	semaphoreS := make(chan struct{}, nSG)

	chEmails := make(chan entities.Email, len(paths))

	wRg.Add(1)
	go func() {
		for _, path := range paths {
			wRg.Add(1)
			go ReadEmail(path, &wRg, chEmails, semaphoreR)
		}
		paths = []string{}
		wRg.Done()
	}()
	fmt.Println("Reading emails")

	go func() {
		var batch []entities.Email
		for email := range chEmails {
			batch = append(batch, email)
			if len(batch) >= batchSize {
				wSg.Add(1)
				go SendEmails(url, batch, user, password, &wSg, semaphoreS)
				batch = []entities.Email{}
			}
		}
		wSg.Add(1)
		go SendEmails(url, batch, user, password, &wSg, semaphoreS)
	}()

	wRg.Wait()
	fmt.Println("Stop reading emails")

	wSg.Wait()
	fmt.Println("Stop sending emails")

	close(chEmails)
}
