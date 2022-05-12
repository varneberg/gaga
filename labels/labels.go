package labels

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/varneberg/gaga/requests"
	"log"
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
	NewName     string `json:"new_name"` // Required to be a json array
	Description string `json:"description,omitempty"`
	Color       string `json:"color,omitempty"`
}

func updateLabel() {}

func parseLabelName(labelName string) []byte {
	var body, err = json.Marshal(map[string][]string{
		"labels": toList(labelName),
	})
	if err != nil {
		log.Fatalln(err)
	}
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
	NodeId      string
	Url         string
	Name        string
	Color       string
	Default     bool
	Description string
}

func GetRepoLabels() []labelResp {
	url := requests.GetRepoUrl()
	body := requests.SendRequest("GET", url, nil)
	if body == nil {
		fmt.Println("Unable to fetch labels")
	}
	var resp []labelResp
	jsonErr := json.Unmarshal(body, &resp)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	return resp
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
	//var body, err = json.Marshal(map[string]string{
	//	"labels": labelName,
	//})
	////body, err := json.Marshal(labelName)
	//fmt.Println("Request body: \n", string(body))
	//if err != nil {
	//	log.Fatalln(err)
	//}

	url := requests.GetPRUrl()
	body := parseLabelName(labelName)
	fmt.Println("Api Request Body: ", string(body))
	requests.SendRequest("POST", url, body)
}

func addNewLabelRepo(label newLabel) {

}

func toList(inputString string) []string {
	var out []string
	out = append(out, inputString)
	return out
}

func LabelHandler(args []string) {
	labelFlag := flag.NewFlagSet("label", flag.ExitOnError)
	labelName := labelFlag.String("n", "", "Name new labels to add")
	labelDesc := labelFlag.String("d", "", "Description of labels, enclosed with \"\"")
	var labelColor = labelFlag.String("c", "", "Color of labels")
	labelFlag.Parse(args)

	addLabelPR(*labelName)
	if labelExists(*labelName) {
		fmt.Println("Label", *labelName, "exists")
		return
	}
	newLabel := newLabel{
		Name:        *labelName,
		Description: *labelDesc,
		Color:       *labelColor,
	}
	//addNewLabelRepo(newLabel)
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
