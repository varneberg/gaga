package labels

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/varneberg/gaga/flags"
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

// Returns URL to the active pull request
func GetPRUrl() string {
	prNumber := strings.Split(ghRefName, "/")[0]
	return ghAPIURL + "/repos/" + ghRepo + "/issues/" + prNumber + "/labels"
}

func GetRepoUrl() string {
	// https://api.github.com/repos/OWNER/REPO/labels
	return ghAPIURL + "/repos/" + ghRepo + "/labels"
}

// colors
// orange : #D93F0B
type Label struct {
	Name        []string `json:"labels"`
	Description string   `json:"description,omitempty"`
	Color       string   `json:"color,omitempty"`
}

func parseLabel(label Label) []byte {
	rb, err := json.Marshal(label)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(rb))
	return rb
}

// Function for sending requests to the github API
func APIRequest(requestType string, url string, requestBody []byte) []byte {
	timeout := time.Duration(5 * time.Second)
	client := &http.Client{
		Timeout: timeout,
	}
	//request, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	request, err := http.NewRequest(requestType, url, bytes.NewBuffer(requestBody))
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

	fmt.Println("Api Response: ", resp.Status)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
		os.Exit(2)
	}
	//fmt.Printf(string(body))
	//fmt.Println()
	return body
}

type labelResp struct {
	ID          int
	Node_id     string
	Url         string
	Name        string
	Color       string
	Default     bool
	Description string
}

//{
//	"id":4093697616,
//	"node_id":"LA_kwDOHRgjdM70AN5Q",
//	"url":"https://api.github.com/repos/varneberg/gaga/labels/bug",
//	"name":"bug",
//	"color":"d73a4a",
//	"default":true,
//	"description":"Something isn't working"
//}
//APIRequest("GET")

func TestGetRepoLabels() []labelResp {
	demoJson := `[
		{"id": 4093697616,"node_id": "LA_kwDOHRgjdM70AN5Q","url": "https://api.github.com/repos/varneberg/gaga/labels/bug","name": "bug","color": "d73a4a","Default": true,"description":"Something isn't working"},
		{"id":4115856884,"node_id":"LA_kwDOHRgjdM71Uv30","url":"https://api.github.com/repos/varneberg/gaga/labels/dependencies","name":"dependencies","color":"0366d6","default":false,"description":"Pull requests that update a dependency file"},
		{"id":4093697619,"node_id":"LA_kwDOHRgjdM70AN5T","url":"https://api.github.com/repos/varneberg/gaga/labels/documentation","name":"documentation","color":"0075ca","default":true,"description":"Improvements or additions to documentation"},
		{"id":4093697621,"node_id":"LA_kwDOHRgjdM70AN5V","url":"https://api.github.com/repos/varneberg/gaga/labels/duplicate","name":"duplicate","color":"cfd3d7","default":true,"description":"This issue or pull request already exists"},
		{"id":4093697623,"node_id":"LA_kwDOHRgjdM70AN5X","url":"https://api.github.com/repos/varneberg/gaga/labels/enhancement","name":"enhancement","color":"a2eeef","default":true,"description":"New feature or request"},
		{"id":4115856887,"node_id":"LA_kwDOHRgjdM71Uv33","url":"https://api.github.com/repos/varneberg/gaga/labels/github_actions","name":"github_actions","color":"000000","default":false,"description":"Pull requests that update GitHub Actions code"},
		{"id":4093697626,"node_id":"LA_kwDOHRgjdM70AN5a","url":"https://api.github.com/repos/varneberg/gaga/labels/good%!f(MISSING)irst%!i(MISSING)ssue","name":"good first issue","color":"7057ff","default":true,"description":"Good for newcomers"},
		{"id":4093697624,"node_id":"LA_kwDOHRgjdM70AN5Y","url":"https://api.github.com/repos/varneberg/gaga/labels/help%!w(MISSING)anted","name":"help wanted","color":"008672","default":true,"description":"Extra attention is needed"},
		{"id":4093697627,"node_id":"LA_kwDOHRgjdM70AN5b","url":"https://api.github.com/repos/varneberg/gaga/labels/invalid","name":"invalid","color":"e4e669","default":true,"description":"This doesn't seem right"}
		]`
	//jsonErr := json.Unmarshal([]byte(demoJson), &resp)
	var resp []labelResp

	jsonErr := json.Unmarshal([]byte(demoJson), &resp)

	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	//fmt.Println(resp)
	//for _, i := range resp {
	//	fmt.Println(i.Name)
	//}
	return resp
}

func GetRepoLabels() []labelResp {
	url := GetRepoUrl()
	body := APIRequest("GET", url, nil)
	var resp []labelResp
	jsonErr := json.Unmarshal(body, &resp)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	return resp
}

func labelExists(labelName string) bool {
	labels := TestGetRepoLabels()
	for _, i := range labels {
		fmt.Println(i)
	}
	return true
}

// Adds labels to current pull request
func addLabelPR(label Label) {
	requestBody := parseLabel(label)
	url := GetPRUrl()
	APIRequest("POST", url, requestBody)
}

func createNewLabelRepo(label Label) {

}

func LabelHandler(args []string) {
	TestGetRepoLabels()
	labelExists("test")

	var labelName flags.FlagSlice
	labelFlag := flag.NewFlagSet("label", flag.ExitOnError)
	labelFlag.Var(&labelName, "n", "Name of the labels")
	//labelNewName := labelFlag.String("N", "", "Name new labels to add")
	labelDesc := labelFlag.String("d", "", "Description of labels, enclosed with \"\"")
	var labelColor = labelFlag.String("c", "", "Color of labels")
	labelFlag.Parse(args)

	//tail := flag.Args()
	//fmt.Printf("Tail: %+q\n", tail)

	newLabel := Label{
		Name:        labelName,
		Description: *labelDesc,
		Color:       *labelColor,
	}
	if labelDesc == nil {
		fmt.Println("No desc")
	}
	if labelColor == nil {
		fmt.Println("No color")
	}

	//if newLabel.Description == "" {
	//	fmt.Println("No description")
	//}

	//fmt.Println("labelName: ", labelName)
	////fmt.Println("labelNewName: ", *labelNewName)
	//fmt.Println("labelDesc: ", *labelDesc)
	//fmt.Println("labelColor: ", *labelColor)
	//fmt.Println()
	fmt.Println(newLabel)
	//addLabelPR(newLabel)

}
