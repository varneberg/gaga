package labels

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/varneberg/gaga/requests"
	"log"
	"os"
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

// colors
// orange : #D93F0B

// New label object
type newLabel struct {
	Name        string `json:"labels"` // Required to be a json array
	Description string `json:"description,omitempty"`
	Color       string `json:"color,omitempty"`
}

type changeLabel struct {
	New_Name    string `json:"new_name"` // Required to be a json array
	Description string `json:"description,omitempty"`
	Color       string `json:"color,omitempty"`
}

func updateLabel() {}

func parseLabelName(labelName string) []byte {
	body, err := json.Marshal(labelName)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(body))
	return body
}

func parseNewLabel(label newLabel) []byte {
	body, err := json.Marshal(label)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(body))
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
	return resp
}

func GetRepoLabels() []labelResp {
	url := requests.GetRepoUrl()
	body := requests.SendRequest("GET", url, nil)
	if body == nil {
		fmt.Println("Ouf")
	}
	var resp []labelResp
	jsonErr := json.Unmarshal(body, &resp)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	return resp
}

func TestlabelExists(labelName string) bool {
	labels := TestGetRepoLabels()
	for _, elem := range labels {
		if labelName == elem.Name {
			return true
		}
	}
	return false
}

func labelExists(labelName string) bool {
	labels := GetRepoLabels()
	for _, elem := range labels {
		if labelName == elem.Name {
			return true
		}
	}
	return false
}

// Adds labels to current pull request
func addLabelPR(labelName string) {
	//requestBody := parseNewLabel(labelName)
	body, err := json.Marshal(labelName)
	if err != nil {
		log.Fatalln(err)
	}
	url := requests.GetPRUrl()
	requests.SendRequest("POST", url, body)
}

func addNewLabelRepo(label newLabel) {

}

func LabelHandler(args []string) {
	//GetRepoLabels()
	//labelExists("test")
	//var labelName flags.FlagSlice
	labelFlag := flag.NewFlagSet("label", flag.ExitOnError)
	//labelFlag.Var(&labelName, "n", "Name of the labels")
	labelName := labelFlag.String("n", "", "Name new labels to add")
	labelDesc := labelFlag.String("d", "", "Description of labels, enclosed with \"\"")
	var labelColor = labelFlag.String("c", "", "Color of labels")
	labelFlag.Parse(args)
	if labelExists(*labelName) {
		fmt.Println("Label", *labelName, "exists")
		fmt.Println()
		return
	}
	newLabel := newLabel{
		Name:        *labelName,
		Description: *labelDesc,
		Color:       *labelColor,
	}
	addNewLabelRepo(newLabel)
	fmt.Println("newLabel: ", newLabel)

	//tail := flag.Args()
	//fmt.Printf("Tail: %+q\n", tail)

	//if newLabel.Description == "" {
	//	fmt.Println("No description")
	//}

	//fmt.Println("labelName: ", labelName)
	////fmt.Println("labelNewName: ", *labelNewName)
	//fmt.Println("labelDesc: ", *labelDesc)
	//fmt.Println("labelColor: ", *labelColor)
	//fmt.Println()
	//addLabelPR(newLabel)

}
