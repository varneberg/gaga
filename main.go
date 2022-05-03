package main

import (
	"fmt"
	"log"
	"os"
)


func main() {
	token := os.Getenv("GTIHUB_TOKEN")
	if token == "" {
		log.Fatal("Error: No Github token present")
	}
	fmt.Printf(token)
}
