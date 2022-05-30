package labels

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
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
	//body := requestBodyuests.ResponseBody(response)
	//requests.CloseRequest(response)

	var lresp []labelResp
	jsonErr := json.Unmarshal(body, &lresp)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	return lresp
}

// Check if label already exist in repository
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
	url := requests.GetPRUrl()
	body := parseLabelName(labelName)
	fmt.Println("Api Request Body: ", body)
	status, resp := requests.SendRequest("POST", url, body)
	//fmt.Println(requests.ResponseStatus(response))
	//fmt.Println(string(requests.ResponseBody(response)))
	//requests.CloseRequest(response)
	fmt.Println(status, "\n", string(resp))
}

func toList(inputString string) []string {
	var out []string
	out = append(out, inputString)
	return out
}

func removeLabel(labelName string) {
	url := requests.GetPRUrl()
	//var body []byte
	body := parseLabelName(labelName)
	status, body := requests.SendRequest("DELETE", url, body)
	fmt.Println(status, "\n", string(body))
}

// Remove all labels from a pull request
func removeAllLabels() {
	url := requests.GetPRUrl()
	//var body []byte
	status, body := requests.SendRequest("DELETE", url, nil)
	fmt.Println(status, "\n", string(body))
	fmt.Println()
}

func createNewLabel(label newLabel) {
	if labelExists(label.Name) {

	}
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
	//Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println(strings.Join(args, " "))
		//LabelHandler(strings.Join(args, " "))
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

	// Check if label already exists in repo
	//if labelExists(labelName) {
	//	fmt.Println("Label", labelName, "exists")
	//}

	// If color nor description is specified

	if labelColor == "" && labelDescription == "" {
		//fmt.Println("Color and description not set")
		addLabelPR(labelName)
		return
	}
	newLabel := newLabel{
		Name:        labelName,
		Description: labelDescription,
		Color:       labelColor,
	}
	createNewLabel(newLabel)
	fmt.Println("newLabel: ", newLabel)
	//addNewLabelRepo(newLabel)

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
