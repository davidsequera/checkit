package main

import (
	"fmt"
	"indexer/controllers"
	"log"
	"os"
	"runtime/pprof"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}
	fmt.Println(os.Getenv("ZINC_FIRST_ADMIN_USER"))

	f, err := os.Create("cpu.prof")
	if err != nil {
		log.Fatal(err)
	}
	pprof.StartCPUProfile(f)
	controllers.Manage("examples/big/fischer-m", "http://localhost:4080/api/_bulkv2")
	pprof.StopCPUProfile()
}
