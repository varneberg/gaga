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

func parseLabel(label string) []byte {
	var labels []string
	labels = append(labels, label)
	rb, err := json.Marshal(map[string][]string{
		"labels": labels,
	})
	if err != nil {
		log.Fatalln(err)
	}
	return rb
}

func postLabel(label string) {
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

type Label struct {
	Name 					string 	`json:"name`
	Description 	string 	`json:description`
	Color					string  `json:color`
}

func formatLabel(label Label){
	fmt.Println(label)

}

func main() {
	checkArgs()
	checkEnv()
	fmt.Println()

	labelcmd := flag.NewFlagSet("label", flag.ExitOnError)
	label_name := labelcmd.String("n", "", "Name of the new label")
	label_desc := labelcmd.String("d", "", "Description of label")
	label_color := labelcmd.String("c", "", "Color of label")
	switch os.Args[1] {
	case "label":
		labelcmd.Parse(os.Args[2:])
		// fmt.Println("Label:")
		// fmt.Println("\tLabel Name: ", *label_name)
		// fmt.Println("\tDescription: ", *label_desc)
		// fmt.Println("\tColor: ", *label_color)
		newLabel := Label{
			Name: *label_name,
			Description: *label_desc,
			Color: *label_color,
		}
		formatLabel(newLabel)
		
	default:
		fmt.Println("what")
	}


	// var options []string = nil
	// for i := 1; i < len(os.Args); i++ {
	// 	fmt.Println("i: ",i)
	// 	switch os.Args[i] {
	// 	case "-l":
	// 		options = nil
	// 		continue
	// 		//postLabel(os.Args[i+1])
	// 	case "-c":
	// 		options = nil
	// 		continue
	// 		//Todo
	// 	default:
	// 		options = append(options, os.Args[i])
	// 		//Todo
	// 	}
	// 	fmt.Println(options)
	// }
}

// colors
// orange : #D93F0B