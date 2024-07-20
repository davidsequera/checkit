package main

import (
	"fmt"
	"indexer/controllers"
	"log"
	"os"
	"runtime"
	"runtime/pprof"

	"github.com/joho/godotenv"
)

func main() {

	args := os.Args

	if len(args) < 2 {
		fmt.Println("No argument. Please insert the direction for indexing")
		return
	}

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}
	user, password := os.Getenv("ZINC_FIRST_ADMIN_USER"), os.Getenv("ZINC_FIRST_ADMIN_PASSWORD")

	f, err := os.Create("./profile/big/cpu.prof")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	runtime.GC()
	mf, err := os.Create("./profile/big/mem.prof")
	if err != nil {
		log.Fatal(err)
	}
	defer mf.Close()
	pprof.WriteHeapProfile(mf)
	pprof.StartCPUProfile(f)
	controllers.Manage(args[1], "http://localhost:4080/api/_bulkv2", &user, &password)
	pprof.StopCPUProfile()

}
