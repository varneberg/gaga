package main

import (
	"bytes"
	"encoding/json"
	// "flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

var (
	ghRepoOwner                  = os.Getenv("GITHUB_REPOSITORY_OWNER")
	ghRef                        = os.Getenv("GITHUB_REF")
	ghRefName                    = os.Getenv("GITHUB_REF_NAME")
	ghRepo                       = os.Getenv("GITHUB_REPOSITORY")
	ghToken                      = os.Getenv("GITHUB_TOKEN")
	ghEvent                      = os.Getenv("GITHUB_EVENT_NAME")
	ghActor                      = os.Getenv("GITHUB_ACTOR")
	ghWorkflow                   = os.Getenv("GITHUB_WORKFLOW")
	ghActionsIDTokenRequestURL   = os.Getenv("ACTIONS_ID_TOKEN_REQUEST_URL")
	ghActionsIDTokenRequestToken = os.Getenv("ACTIONS_ID_TOKEN_REQUEST_TOKEN")
	ghAPIURL                     = os.Getenv("GITHUB_API_URL")
)

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

func inputLabels() []string {
	input := os.Args[1:]
	return input
}

func checkEnv() {
	if ghEvent != "pull_request" {
		fmt.Println("Error: Not a pull request")
		os.Exit(0)
	}
	if ghRefName == "" {
		fmt.Println("Error: Github Reference Name not available")
		os.Exit(0)
	}
}

func apiRequest() {
	return
}

func postLabel(labels []string) {
	//inputLabels()

	// labels := []string{"test", "test2"}
	// labels := inputLabels()

	requestBody, err := json.Marshal(map[string][]string{
		"labels": labels,
	})
	// fmt.Println(string(requestBody))

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


var labels []string


func main() {
	if len(os.Args[1:]) == 0 {
		fmt.Println("Error: No arguments")
		os.Exit(0)
	}
	
	checkEnv()
	for i, arg := range os.Args[1:] {

		if i+2 >= len(os.Args){
			break
		}
		value := os.Args[i+2]
		fmt.Println(arg, value)
		switch arg {
		case "-l":
			labels[i] = value
		case "-c":
			// comment := value
		}

		}
		
		postLabel(labels)



}
