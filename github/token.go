package github

import (
	"fmt"
	"log"
	"os"
)

func getToken() { 
	token := os.Getenv("GTIHUB_AUTH_TOKEN")
	if token == "" {
		log.Fatal("Error: No Github token present")
	}
	fmt.Println(token)
}