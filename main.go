package main

import (
	"bytes"
	"encoding/json"
	"flag"
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

func checkArgs() {
	if len(os.Args[1:]) == 0 {
		fmt.Println("Error: No arguments")
		os.Exit(0)
	}
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


func postLabel(label Label) {
	requestBody := parseLabel(label)
	prNumber := strings.Split(ghRefName, "/")[0]
	url := ghAPIURL + "/repos/" + ghRepo + "/issues/" + prNumber + "/labels"
	fmt.Println("URL: ", url)

	timeout := time.Duration(5 * time.Second)
	client := &http.Client{
		Timeout: timeout,
	}

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
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
	fmt.Println()
}

type arrayFlags []string

func(i *arrayFlags) String() string {
	return ""

}

func(i *arrayFlags) Set(value string) error {
	*i = append(*i, strings.TrimSpace(value))
	return nil
}
// colors
// orange : #D93F0B
type Label struct {
	Name 					[]string 	`json:"name"`
	Description 	string 	`json:"description,omitempty"`
	Color					string  `json:"color,omitempty"`
}

func parseLabel(label Label) []byte {
	rb, err := json.Marshal(label)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(rb))
	return rb
}

func main() {
	checkArgs()
	checkEnv()
	// var label_name arrayFlags
	labelcmd := flag.NewFlagSet("label", flag.ExitOnError)
	label_name := labelcmd.String("n", "", "Name of the new label")
	// labelcmd.Var(&label_name,"n", "Name of the label")
	label_desc := labelcmd.String("d", "", "Description of label, enclosed with \"\"")
	label_color := labelcmd.String("c", "", "Color of label")
	switch os.Args[1] {
	case "label":
		var labelList []string
		labelcmd.Parse(os.Args[2:])
		labelList = append(labelList, *label_name)
		newLabel := Label{
			Name: labelList,
			Description: *label_desc,
			Color: *label_color,
		}
		postLabel(newLabel)
		
	default:
		fmt.Println("Invalid arguments")
	}
}
