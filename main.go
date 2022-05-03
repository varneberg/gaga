package main

import (
	"context"
	"fmt"
	"log"
	"os"
	//"github.com/google/go-github/github"
	//"golang.org/x/oauth2"
)


func main() {
	token := os.Getenv("$GTIHUB_TOKEN")
	//ctx := context.Background()
	if token == ""{
		log.Fatal("Unauthorized: No token present")
	} else {
		fmt.Println("Success: Token aquired!")
	}
}
