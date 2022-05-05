package main

import (
	// "bytes"
	// "encoding/json"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	// "log"
	"net/http"
	"os"
	"strings"
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

func auth() {
	if ghEvent != "pull_request" {fmt.Println("Error: Not a pull request")}
	if ghRefName == "" {fmt.Println("Error: Github Reference Name not available")}

	requestBody, err := json.Marshal(map[string]string{
		"labels": "test",
	})
	if err != nil {log.Fatalln(err)}

	prNumber := strings.Split(ghRefName, "/")[0]
	url := ghAPIURL + "/repos/" + ghRepo + "/issues/" + prNumber + "/labels"
	fmt.Println("URL: ", url)

	timeout := time.Duration(5 * time.Second)
	client := &http.Client{
		Timeout: timeout,
	}

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	request.Header.Add("Authorization", fmt.Sprintf("token %d", ghToken))
	request.Header.Set("Content-Type", "application/json")
	if err != nil {log.Fatalln(err)}
	
	resp, err := client.Do(request)
	if err != nil {log.Fatalln(err)}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {log.Fatalln(err)}

	log.Println(string(body))
	// req, err := http.NewRequest("POST", url, nil)
	// if err != nil {fmt.Println("Error: ", err.Error())}
	// req.Header.Add("Authorization: Bearer", ghToken)
	// req.Header.Add("Content-Type", "application/json")
	// fmt.Print(req)
	// response, err := client.Do(req)
	// if err != nil {fmt.Println("Error: ",err.Error())}
	// defer response.Body.Close()
}

func main() {
	listEnv()
	fmt.Println()
	auth()
}

// func getRequest(){
// 	resp, err := http.Get("https://www.vg.no")
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	defer resp.Body.Close()

// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	log.Println(string(body))
// }
// func auth(){
// 	ctx := context.Background()
// 	if ghToken == ""{
// 		fmt.Println("Error: Github token is missing")
// 		return
// 	}
// 	ts := oauth2.StaticTokenSource(
// 		&oauth2.Token{AccessToken: ghToken},
// 	)

// 	tc := oauth2.NewClient(ctx, ts)

// 	client := github.NewClient(tc)
// 	user, resp, err := client.Users.Get(ctx, "")
// 	if err != nil {
// 		fmt.Printf("\nerror: %v\n", err)
// 		return
// 	}
// 	// Rate.Limit should most likely be 5000 when authorized.
// 	log.Printf("Rate: %#v\n", resp.Rate)
// 	// If a Token Expiration has been set, it will be displayed.
// 	fmt.Printf("\n%v\n", github.Stringify(user))

// }
// func main() {
//fmt.Print("GitHub Token: ")

// for _, env := range os.Environ() {
// 	fmt.Println(env)
// }
// byteToken, _ := term.ReadPassword(int(syscall.Stdin))
// println()
// token := string(byteToken)

// ctx := context.Background()
// ts := oauth2.StaticTokenSource(
// 	&oauth2.Token{AccessToken: token},
// )
// tc := oauth2.NewClient(ctx, ts)

// client := github.NewClient(tc)

// user, resp, err := client.Users.Get(ctx, "")
// if err != nil {
// 	fmt.Printf("\nerror: %v\n", err)
// 	return
// }

// // Rate.Limit should most likely be 5000 when authorized.
// log.Printf("Rate: %#v\n", resp.Rate)

// // If a Token Expiration has been set, it will be displayed.
// if !resp.TokenExpiration.IsZero() {
// 	log.Printf("Token Expiration: %v\n", resp.TokenExpiration)
// }

// fmt.Printf("\n%v\n", github.Stringify(user))
// }
