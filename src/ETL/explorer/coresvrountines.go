package main

import (
	"fmt"
	"runtime"
)

func main2() {
	numCPU := runtime.NumCPU() // Get number of CPU cores
	fmt.Printf("Number of CPU cores: %d\n", numCPU)

	// Set GOMAXPROCS to use all CPU cores
	prevGOMAXPROCS := runtime.GOMAXPROCS(numCPU)
	fmt.Printf("Previous value of GOMAXPROCS: %d\n", prevGOMAXPROCS)

	// Additional work to demonstrate parallelism
	for i := 0; i < numCPU; i++ {
		go func() {
			fmt.Println("Hello from a goroutine!")
		}()
	}

	// Wait for goroutines to finish
	fmt.Println("Waiting for goroutines to finish...")
	fmt.Scanln() // Wait for enter key press
}
