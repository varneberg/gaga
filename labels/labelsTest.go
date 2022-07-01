package labels

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/varneberg/gaga/requests"
)

// colors
// orange : #D93F0B

func TestupdateLabel() {}

func TestGetRepoLabels() []labelResp {
	demoJson := `[
		{"id": 4093697616,"node_id": "LA_kwDOHRgjdM70AN5Q","url": "https://api.github.com/repos/varneberg/gaga/labels/bug","name": "bug","color": "d73a4a","Default": true,"description":"Something isn't working"},
		{"id":4115856884,"node_id":"LA_kwDOHRgjdM71Uv30","url":"https://api.github.com/repos/varneberg/gaga/labels/dependencies","name":"dependencies","color":"0366d6","default":false,"description":"Pull requests that update a dependency file"},
		{"id":4093697619,"node_id":"LA_kwDOHRgjdM70AN5T","url":"https://api.github.com/repos/varneberg/gaga/labels/documentation","name":"documentation","color":"0075ca","default":true,"description":"Improvements or additions to documentation"},
		{"id":4093697621,"node_id":"LA_kwDOHRgjdM70AN5V","url":"https://api.github.com/repos/varneberg/gaga/labels/duplicate","name":"duplicate","color":"cfd3d7","default":true,"description":"This issue or pull request already exists"},
		{"id":4093697623,"node_id":"LA_kwDOHRgjdM70AN5X","url":"https://api.github.com/repos/varneberg/gaga/labels/enhancement","name":"enhancement","color":"a2eeef","default":true,"description":"New feature or request"},
		{"id":4115856887,"node_id":"LA_kwDOHRgjdM71Uv33","url":"https://api.github.com/repos/varneberg/gaga/labels/github_actions","name":"github_actions","color":"000000","default":false,"description":"Pull requests that update GitHub Actions code"},
		{"id":4093697624,"node_id":"LA_kwDOHRgjdM70AN5Y","url":"https://api.github.com/repos/varneberg/gaga/labels/help wanted","name":"help wanted","color":"008672","default":true,"description":"Extra attention is needed"},
		{"id":4093697627,"node_id":"LA_kwDOHRgjdM70AN5b","url":"https://api.github.com/repos/varneberg/gaga/labels/invalid","name":"invalid","color":"e4e669","default":true,"description":"This doesn't seem right"}
		]`
	//jsonErr := json.Unmarshal([]byte(demoJson), &resp)
	var resp []labelResp
	jsonErr := json.Unmarshal([]byte(demoJson), &resp)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	fmt.Println("Existing test labels: ")
	for _, i := range resp {
		fmt.Println("\t", i)
	}
	return resp
}

func TestLabelExists(labelName string) bool {
	labels := TestGetRepoLabels()
	for _, elem := range labels {
		if labelName == elem.Name {
			return true
		}
	}
	return false
}

func TestAddLabelPR(labelName string) {
	url := requests.GetLabelUrl()
	body := parseLabelName(labelName)
	fmt.Println("Api Request body: \n\t", string(body))
	requests.TestSendRequest("POST", url, body)
}

func TestLabelHandler() {

	if TestLabelExists(labelName) {
		fmt.Println("Label", labelName, "exists")
	}

	// If color nor description is specified
	if labelColor == "" && labelDescription == "" {
		//fmt.Println("Color and description not set")
		TestAddLabelPR(labelName)
		return
	}
	newLabel := newLabel{
		Name:        labelName,
		Description: labelDescription,
		Color:       labelColor,
	}
	//addNewLabelRepo(newLabel)
	fmt.Println("newLabel: ", newLabel)
}
