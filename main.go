package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

var ghRepoOwner = os.Getenv("GITHUB_REPOSITORY_OWNER")
var ghRef = os.Getenv("GITHUB_REF")
var ghRefName = os.Getenv("GITHUB_REF_NAME")
var ghRepo = os.Getenv("GITHUB_REPOSITORY")
var ghToken = os.Getenv("GITHUB_TOKEN")
var ghEvent = os.Getenv("GITHUB_EVENT_NAME")
var ghActor = os.Getenv("GITHUB_ACTOR")
var ghWorkflow = os.Getenv("GITHUB_WORKFLOW")
var ghActionsIDTokenRequestURL = os.Getenv("ACTIONS_ID_TOKEN_REQUEST_URL")
var ghActionsIDTokenRequestToken = os.Getenv("ACTIONS_ID_TOKEN_REQUEST_TOKEN")
var ghAPIURL = os.Getenv("GITHUB_API_URL")

func listEnv() {
	fmt.Println("Github Repo Owner: ", ghRepoOwner)
	fmt.Println("Github Actor: ", ghActor)
	fmt.Println("Github Ref: ", ghRef)
	fmt.Println("Github Ref Name: ", ghRefName)
	fmt.Println("Github API URL: ", ghAPIURL)
	fmt.Println("Github Actions Token Request URL: ", ghActionsIDTokenRequestURL)
	fmt.Println("Github Repo: ", ghRepo)
	fmt.Println("Github Event: ", ghEvent)
	fmt.Println("Github Workflow: ", ghWorkflow)
	fmt.Println("Github Token: ", ghToken)
	fmt.Println("Github Actions Request Token: ", ghActionsIDTokenRequestToken)
	fmt.Printf("-----------------------\n")
}

func inputLabels() ([]string){
	input := os.Args[1:]
	// fmt.Println(input)
	return input
}

func auth() {
	inputLabels()
	if ghEvent != "pull_request" {
		fmt.Println("Error: Not a pull request")
	}
	if ghRefName == "" {
		fmt.Println("Error: Github Reference Name not available")
	}

	// labels := []string{"test", "test2"}
	labels := inputLabels()

	requestBody, err := json.Marshal(map[string][]string{
		"labels": labels,
	})
	fmt.Println(string(requestBody))
	if err != nil {
		log.Fatalln(err)
	}

	prNumber := strings.Split(ghRefName, "/")[0]
	url := ghAPIURL + "/repos/" + ghRepo + "/issues/" + prNumber + "/labels"
	fmt.Println("URL: ", url)

	timeout := time.Duration(5 * time.Second)
	client := &http.Client{
		Timeout: timeout,
	}

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	// request, err := http.NewRequest("POST", url, nil)
	request.Header.Add("Accept", "application/vnd.github.v3+json")
	request.Header.Add("Authorization", "token "+ghToken)
	request.Header.Set("Content-Type", "application/json")
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := client.Do(request)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf(string(body))
}

func main() {
	listEnv()
	fmt.Println()
	auth()
}
