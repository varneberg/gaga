package main

import (
	"fmt"
	"log"
	"os"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"

)


func main() {
	token := os.Getenv("GTIHUB_TOKEN")
	repo := os.Getenv("GITHUB_REPOSITORY_NAME")
	owner := os.Getenv("GITHUB_REPOSITY_NAME")
	fmt.Printf("Github Token:",token)
}
