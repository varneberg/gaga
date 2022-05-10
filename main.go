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


func checkArgs() {
	if len(os.Args[1:]) == 0 {
		fmt.Println("Error: No arguments")
		os.Exit(0)
	}
}

func checkEnv() {
	if ghEvent != "pull_request" {
		fmt.Println("Error: Not a pull request")
		os.Exit(2)
	}
	if ghRefName == "" {
		fmt.Println("Error: Github Reference Name not available")
		os.Exit(2)
	}
}

// Returns URL to the active pull request
func getPrUrl() string {
	prNumber := strings.Split(ghRefName, "/")[0]
	return ghAPIURL + "/repos/" + ghRepo + "/issues/" + prNumber + "/labels"
}

func getRepoUrl() string{
	// https://api.github.com/repos/OWNER/REPO/labels
	return ghAPIURL + "/repos/" + ghRepo + "/labels"
}

type flagSlice []string

func(i *flagSlice) String() string {return ""}


func(i *flagSlice) Set(value string) error {
	*i = append(*i, strings.TrimSpace(value))
	return nil
}
// colors
// orange : #D93F0B
type Label struct {
	Name 					[]string 	  `json:"labels"`
	Description 	  string 		`json:"description,omitempty"`
	Color					  string  	`json:"color,omitempty"`
}

func parseLabel(label Label) []byte {
	rb, err := json.Marshal(label)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(rb))
	return rb
}

func apiRequest(requestBody []byte, url string, ){
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
		os.Exit(2)
	}
	fmt.Printf(string(body))
	fmt.Println()
}

// Adds labels to current pull request
func addLabelPR(label Label) {
	requestBody := parseLabel(label)
	url := getPrUrl()
	apiRequest(requestBody, url)
	// requestBody := parseLabel(label)
	// prNumber := strings.Split(ghRefName, "/")[0]
	// url := ghAPIURL + "/repos/" + ghRepo + "/issues/" + prNumber + "/labels"
	// url := getPrUrl()
	// fmt.Println("PR_URL: ", url)

	// timeout := time.Duration(5 * time.Second)
	// client := &http.Client{
	// 	Timeout: timeout,
	// }

	// request, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	// request.Header.Add("Accept", "application/vnd.github.v3+json")
	// request.Header.Add("Authorization", "token "+ghToken)
	// request.Header.Set("Content-Type", "application/json")
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// resp, err := client.Do(request)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// defer resp.Body.Close()

	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// fmt.Printf(string(body))
	// fmt.Println()
}

func addLabelRepo(label Label){

}

func isFlagPassed(name string) bool {
    found := false
    flag.Visit(func(f *flag.Flag) {
        if f.Name == name {
            found = true
        }
    })
    return found
}
func main() {
	checkArgs()
	checkEnv()
	var label_name flagSlice
	labelcmd := flag.NewFlagSet("label", flag.ExitOnError)
	// label_name := labelcmd.String("n", "", "Name of the label")
	labelcmd.Var(&label_name,"n", "Name of the label")
	// label_newLabel := labelcmd.String("N", "", "Name new label to add")
	label_desc := labelcmd.String("d", "", "Description of label, enclosed with \"\"")
	label_color := labelcmd.String("c", "", "Color of label")
	switch os.Args[1] {
	case "label":
		// var labelList []string
		labelcmd.Parse(os.Args[2:])
		// labelList = append(labelList, *label_name)

		fmt.Println(os.Args[2])
		newLabel := Label{
			Name: label_name,
			Description: *label_desc,
			Color: *label_color,
		}
		addLabelPR(newLabel)
		
	default:
		fmt.Println("Invalid arguments")
	}
	
}
