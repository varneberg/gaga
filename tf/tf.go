package tf

import (
	// "fmt"
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
	"github.com/varneberg/gaga/comments"
	"github.com/varneberg/gaga/labels"
	"github.com/varneberg/gaga/parser"
)

var TFCmd = &cobra.Command{
	Use:   "tflabel",
	Short: "Labels from terraform plan",
	Long:  `Add labels based on terraform plan (Default: UNIX pipe).`,
	Run: func(cmd *cobra.Command, args []string) {
		terraformHandler()
	},
}

func getPlanInput() string {
	var tfplan string
	if parser.IsInputFromPipe() {
		tfplan = parser.ReadPipeInput()
	} else {
		tfplan = parser.ReadFileInput(readFile)
	}
	return tfplan
}

func parsePlan(tfplan string) string {
	re, err := regexp.Compile(`Plan: [\w] to add, [\w] to change, [\w] to destroy`)
	if err != nil {
		log.Fatalln(err)
	}
	return re.FindString(tfplan)

}

// Comment terraform plan on pull request
func commentPlan(planResult string){
	mdComment := comments.ToMarkdown("Test title",planResult)
	fmt.Println(mdComment)
	if commentPlanFlag{
		comments.PostComment(planResult)
	}
}

// Read terraform pipe input and add corresponding labels to pull request
func getPlanLabels(tfplan string) {
	// tfplan := getPlanInput()

	if outPipeFlag {
		parser.WritePipeOutput(tfplan)
	}
	newChanges := parsePlan(tfplan)
	if newChanges == "" {
		labels.PostLabelPR(labelNoChanges)
		return
	}
	parsed := strings.Trim(newChanges, "Plan: ")
	split := strings.Split(parsed, ", ")
	//fmt.Println(split)
	for _, s := range split {
		diff := string(s[0])
		action := strings.Split(s, " ")[2]
		if diff != "0" { // Check if there are more than 0 changes
			switch action {
			case "add":
				labels.PostLabelPR(labelAddUpdate)
			case "destroy":
				labels.PostLabelPR(labelDestroy)
			case "error":
				labels.PostLabelPR(labelError)
			}
		}
	}
	// commentPlan(tfplan)
}


var labelAddUpdate string
var labelDestroy string
var labelNoChanges string
var labelError string

// var readString string
var readFile string
var outPipeFlag bool
var commentPlanFlag bool

func init() {
	// Customizeable labels based on output from terraform plan, with defaults
	TFCmd.Flags().StringVarP(&labelAddUpdate, "label-add", "a", "tf/add-update", "Terraform Add/Update label name")
	TFCmd.Flags().StringVarP(&labelDestroy, "label-destroy", "d", "tf/destroy", "Terraform destroy label name")
	TFCmd.Flags().StringVarP(&labelNoChanges, "label-no-changes", "n", "tf/no-changes", "Terraform no changes label name")
	TFCmd.Flags().StringVarP(&labelError, "label-error", "e", "tf/error", "Terraform error label name")

	// Terraform plan input source
	TFCmd.Flags().StringVarP(&readFile, "from-file", "f", "", "Read Terraform plan from file") // TODO
	// TFCmd.Flags().StringVarP(&readString, "from-string", "s", "", "Read Terraform plan from string")

	// Terraform plan output
	TFCmd.Flags().BoolVarP(&outPipeFlag, "out", "o", false, "Output Terraform plan as std.out")
	TFCmd.Flags().BoolVarP(&commentPlanFlag, "comment", "c", false, "Comment plan result on pull request")
}

func terraformHandler() {
	tfplan := getPlanInput()
	commentPlan(tfplan)
	getPlanLabels(tfplan)
}
