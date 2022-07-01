package labels

import (
	"encoding/json"
	// "fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/varneberg/gaga/requests"
)

// colors
// orange : #D93F0B

// New label object
type newLabel struct {
	Name        string `json:"name"` // Required to be a json array
	Description string `json:"description,omitempty"`
	Color       string `json:"color,omitempty"`
}

func parseLabelName(labelName string) []byte {
	var body, err = json.Marshal(map[string][]string{
		"labels": toList(labelName),
	})
	if err != nil {
		log.Fatalln(err)
	}
	return body
}

// Rest api response
type labelResp struct {
	ID          int
	NodeId      string
	Url         string
	Name        string
	Color       string
	Default     bool
	Description string
}

// GetRepoLabels Get all labels defined within a repository
func GetRepoLabels() []labelResp {
	url := requests.GetRepoUrl()
	_, body := requests.SendRequest("GET", url, nil)

	var lresp []labelResp
	jsonErr := json.Unmarshal(body, &lresp)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	return lresp
}

// Check if label already exist in repository
func isLabel(labelName string) bool {
	labels := GetRepoLabels()
	for _, elem := range labels {
		if labelName == elem.Name {
			return true
		}
	}
	return false
}

// Adds labels to current pull request
func PostLabelPR(labelName string) {
	url := requests.GetLabelUrl()
	body := parseLabelName(labelName)
	//fmt.Println("Api Request Body: ", string(body))
	status, resp := requests.SendRequest("POST", url, body)
	// requests.PrintResponse(status, resp)
	requests.CheckRespError(status, resp)
}

func toList(inputString string) []string {
	var out []string
	out = append(out, inputString)
	return out
}

func removeLabel(labelName string) {
	url := requests.GetLabelUrl()
	//var body []byte
	body := parseLabelName(labelName)
	status, body := requests.SendRequest("DELETE", url, body)
	// requests.PrintResponse(status, body)
	requests.CheckRespError(status, body)
}

// Remove all labels from a pull request
func removeAllLabels() {
	url := requests.GetLabelUrl()
	status, body := requests.SendRequest("DELETE", url, nil)
	// requests.PrintResponse(status, body)
	requests.CheckRespError(status, body)
}

func createNewLabel(label newLabel) {
	if isLabel(label.Name) {
		PostLabelPR(label.Name)
		return
	}
	url := requests.GetRepoUrl()
	var body, err = json.Marshal(newLabel{
		Name:        label.Name,
		Description: label.Description,
		Color:       label.Color,
	})
	if err != nil {
		log.Fatalln(err)
	}
	status, respbody := requests.SendRequest("POST", url, body)
	// requests.PrintResponse(status, respbody)
	requests.CheckRespError(status, respbody)
	PostLabelPR(label.Name)
}

var labelName string
var labelColor string
var labelDescription string
var labelRemove bool
var removeAll bool

var LabelCmd = &cobra.Command{
	Use:   "label [label]",
	Short: "Label a pull request",
	Long:  `label is for labeling a pull request.`,
	Run: func(cmd *cobra.Command, args []string) {
		LabelHandler()
	},
}

func init() {
	LabelCmd.Flags().StringVarP(&labelName, "name", "n", "", "Label name")
	LabelCmd.Flags().StringVarP(&labelColor, "color", "c", "", "Label color")
	LabelCmd.Flags().StringVarP(&labelDescription, "description", "d", "", "Label description")
	LabelCmd.PersistentFlags().BoolVarP(&labelRemove, "remove", "r", false, "Remove a label")
	LabelCmd.PersistentFlags().BoolVarP(&removeAll, "remove-all", "R", false, "Remove all current labels on PR")
}

func LabelHandler() {

	if labelRemove {
		removeLabel(labelName)
	}
	if removeAll {
		removeAllLabels()
	}

	// If color nor description is specified
	if labelColor == "" && labelDescription == "" {
		PostLabelPR(labelName)
		return
	}
	newLabel := newLabel{
		Name:        labelName,
		Description: labelDescription,
		Color:       labelColor,
	}
	createNewLabel(newLabel)
	//fmt.Println("newLabel: ", newLabel)
}
