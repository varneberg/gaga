package labels

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
	"github.com/varneberg/gaga/parser"
)

var TFCmd = &cobra.Command{
	Use:   "tflabel",
	Short: "Labels from terraform plan",
	Long:  `Add labels based on terraform plan from unix pipe.`,
	//Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		terraformHandler()
	},
}

type tfResult struct {
	Details []string
	Changes string
}

func parseTerraformPlan() tfResult {
	tfplan := parser.ReadPipeInput()
	toSlice := strings.Split(tfplan, "\n")
	var out []string
	for i, line := range toSlice {
		// Separator used in Terraform plan
		if strings.Contains(line, "─────────────────────────────────────────────────────────────────────────────") {
			out = toSlice[i+1 : len(toSlice)-4]
			break
		}
	}
	var result tfResult
	result.Changes = out[len(out)-1]
	result.Details = out[:len(out)-1]
	return result
}

func getPlanResults() {
	tfplan := parser.ReadPipeInput()
	//s := "Plan: 9 to add, 0 to change, 0 to destroy."
	re, err := regexp.Compile(`Plan: [\w] to add, [\w] to change, [\w] to destroy`)
	if err != nil {
		log.Fatalln(err)
	}
	newChanges := re.FindString(tfplan)
	if newChanges == "" {
		AddLabelPR(labelNoChanges)
		return
	}
	parsed := strings.Trim(newChanges, "Plan: ")
	split := strings.Split(parsed, ", ")
	fmt.Println(split)
	for _, s := range split {
		diff := string(s[0])
		action := strings.Split(s, " ")[2]
		if diff != "0" {
			switch action {
			case "add":
				AddLabelPR(labelAddUpdate)
			case "destroy":
				AddLabelPR(labelDestroy)
			case "error":
				AddLabelPR(labelError)

			}
		}

	}

}

func handlePlanResult() {
	//newResult := parseTerraformPlan()
	//fmt.Println(newResult.Changes)
	getPlanResults()

}

var labelAddUpdate string
var labelDestroy string
var labelNoChanges string
var labelError string
var readString string

func init() {
	// Specifyable labels based on output from terraform plan, with defaults
	TFCmd.Flags().StringVarP(&labelAddUpdate, "label-add", "a", "tf/add-update", "Terraform Add/Update label name")
	TFCmd.Flags().StringVarP(&labelDestroy, "label-destroy", "d", "tf/destroy", "Terraform destroy label name")
	TFCmd.Flags().StringVarP(&labelNoChanges, "label-no-changes", "n", "tf/no-changes", "Terraform no changes label name")
	TFCmd.Flags().StringVarP(&labelError, "label-error", "e", "tf/error", "Terraform error label name")

	// Terraform plan input options
	// TFCmd.Flags().BoolVarP(&readPipe, "pipe", "p", true, "Read Terraform plan from pipe")
	TFCmd.Flags().StringVarP(&readString, "from-string", "s", "", "Read Terraform plan from string")
}

func terraformHandler() {
	// planResult := parseTerraformPlan()
	// fmt.Println(planResult)
	handlePlanResult()
}
