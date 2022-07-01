package requests

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

var (
	ghRefName = os.Getenv("GITHUB_REF_NAME")
	ghRepo    = os.Getenv("GITHUB_REPOSITORY")
	ghToken   = os.Getenv("GITHUB_TOKEN")
	ghAPIURL  = os.Getenv("GITHUB_API_URL")
	// ghRepoOwner                  = os.Getenv("GITHUB_REPOSITORY_OWNER")
	// ghRef                        = os.Getenv("GITHUB_REF")
	// ghEvent                      = os.Getenv("GITHUB_EVENT_NAME")
	// ghActor                      = os.Getenv("GITHUB_ACTOR")
	// ghWorkflow                   = os.Getenv("GITHUB_WORKFLOW")
	// ghActionsIDTokenRequestURL   = os.Getenv("ACTIONS_ID_TOKEN_REQUEST_URL")
	// ghActionsIDTokenRequestToken = os.Getenv("ACTIONS_ID_TOKEN_REQUEST_TOKEN")
)

// Get URL of github repository
func GetRepoUrl() string {
	// https://api.github.com/repos/OWNER/REPO/labels
	return ghAPIURL + "/repos/" + ghRepo + "/labels"
}

// Get URL for labeling current pull request
func GetLabelUrl() string {
	prNumber := strings.Split(ghRefName, "/")[0]
	return ghAPIURL + "/repos/" + ghRepo + "/issues/" + prNumber + "/labels"
}

func GetPrURL() string{
	prNumber := strings.Split(ghRefName, "/")[0]
	return ghAPIURL + "/repos/" + ghRepo + "/issues/" + prNumber
}

// SendRequest function for sending requests to the github API
func SendRequest(requestMethod string, url string, requestBody []byte) (int, []byte) {
	timeout := time.Duration(5 * time.Second)
	client := &http.Client{
		Timeout: timeout,
	}
	// Init request and add required headers
	request, err := http.NewRequest(requestMethod, url, bytes.NewBuffer(requestBody))
	request.Header.Add("Accept", "application/vnd.github.v3+json")
	request.Header.Add("Authorization", "token "+ghToken)
	request.Header.Set("Content-Type", "application/json")
	if err != nil {
		log.Fatalln(err)
	}

	// Send request
	resp, err := client.Do(request)
	if err != nil {
		log.Fatalln(err)

	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	statusCode := resp.StatusCode
	return statusCode, body
}

func CheckRespError(respCode int, respBody []byte) {}

func PrintResponse(status int, response []byte) {
	fmt.Println(">> ", status)
	fmt.Println("\t", string(response))
	fmt.Println()

}

func TestSendRequest(requestMethod string, url string, requestBody []byte) {
	request, err := http.NewRequest(requestMethod, url, bytes.NewBuffer(requestBody))
	request.Header.Add("Accept", "application/vnd.github.v3+json")
	request.Header.Add("Authorization", "token "+ghToken)
	request.Header.Set("Content-Type", "application/json")
	if err != nil {
		log.Fatalln(err)
	}

	for _, i := range request.Form {
		fmt.Println(i)

	}

}
