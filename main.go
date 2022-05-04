package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)


func main() {
	ctx := context.Background()
	token := os.Getenv("GTIHUB_TOKEN")

	
	if token == ""{
		log.Fatal("Unauthorized: No token present")
	} else {
		fmt.Println("Success: Token aquired!")
	}
}
