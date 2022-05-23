package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/varneberg/gaga/labels"
	//"github.com/varneberg/gaga/labels"
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

// Check if any input was given
func checkArgs() {
	if len(os.Args[1:]) == 0 {
		fmt.Println("Error: No arguments")
		os.Exit(0)
	}
}

func main() {
	//checkEnv()
	checkArgs()
	fmt.Println()

	//var labelCmd = &cobra.Command{
	//	Use:   "label [label]",
	//	Short: "Label a pull request",
	//	Long:  `label is for labeling a pull request.`,
	//	//Args:  cobra.MinimumNArgs(1),
	//	Run: func(cmd *cobra.Command, args []string) {
	//		fmt.Println(strings.Join(args, " "))
	//	},
	//}
	//
	//var labelName string
	//var labelColor string
	//var labelDescription string
	//labelCmd.Flags().StringVarP(&labelName, "name", "n", "", "label name")
	//labelCmd.Flags().StringVarP(&labelColor, "color", "c", "", "label color")
	//labelCmd.Flags().StringVarP(&labelDescription, "description", "d", "", "label description")

	var rootCmd = &cobra.Command{Use: "app"}
	rootCmd.AddCommand(labels.LabelCmd)
	rootCmd.Execute()
}
