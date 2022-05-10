package flags

import (
	"flag"
	"fmt"
	"os"
	"strings"
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

func checkEnv() {
	if ghEvent != "pull_request" {
		fmt.Println("Error: Not a pull request")
		os.Exit(2)
	}
	if ghRefName == "" {
		fmt.Println("Error: Github Reference Name not available")
		os.Exit(2)
	}
}

func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}

type FlagSlice []string

func (i *FlagSlice) String() string { return "" }

func (i *FlagSlice) Set(value string) error {
	*i = append(*i, strings.TrimSpace(value))
	return nil
}

// Handles input flags
func CmdParser(args []string) {
	//checkArgs()
	//checkEnv()

	//switch args[1] {
	//case "label":
	//	var labelName flagSlice
	//	labelFlag.Var(&labelName, "n", "Name of the labels")
	//	labelNewName := labelFlag.String("N", "", "Name new labels to add")
	//	labelDesc := labelFlag.String("d", "", "Description of labels, enclosed with \"\"")
	//	labelColor := labelFlag.String("c", "", "Color of labels")
	//	labelFlag.Parse(args[2:])
	//	fmt.Println(labelName)
	//	fmt.Println(*labelNewName)
	//	fmt.Println(*labelDesc)
	//	fmt.Println(*labelColor)
	//	labels.LabelHandler()
	//}

	//checkArgs()
	//checkEnv()
	//var label_name flagSlice
	//labelcmd := flag.NewFlagSet("labels", flag.ExitOnError)
	//// label_name := labelcmd.String("n", "", "Name of the labels")
	//labelcmd.Var(&label_name, "n", "Name of the labels")
	//// label_newLabel := labelcmd.String("N", "", "Name new labels to add")
	//label_desc := labelcmd.String("d", "", "Description of labels, enclosed with \"\"")
	//label_color := labelcmd.String("c", "", "Color of labels")
	//switch os.Args[1] {
	//case "labels":
	//	// var labelList []string
	//	labelcmd.Parse(os.Args[2:])
	//	// labelList = append(labelList, *label_name)
	//
	//	fmt.Println(os.Args[2])
	//	newLabel := Label{
	//		Name:        label_name,
	//		Description: *label_desc,
	//		Color:       *label_color,
	//	}
	//	addLabelPR(newLabel)
	//
	//default:
	//	fmt.Println("Invalid arguments")
	//}
}
