package main

import (
	"fmt"
	"github.com/varneberg/gaga/labels"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
)



func checkEnv() {
	if ghEvent != "pull_request" {
		fmt.Println("Error: Not a pull request")
		os.Exit(0)
	}
	if ghRefName == "" {
		fmt.Println("Error: Github Reference Name not available")
)

// Check if any input was given
func checkArgs() {
	if len(os.Args[1:]) == 0 {
		fmt.Println("Error: No arguments")
		os.Exit(0)
	}
}


func parseLabel(label string) []byte {
	var labels []string
	labels = append(labels, label)
	rb, err := json.Marshal(map[string][]string{
		"labels": labels,
	})
	if err != nil {
		log.Fatalln(err)
	}
	return rb
}

func postLabel(label string) {
	requestBody := parseLabel(label)
	prNumber := strings.Split(ghRefName, "/")[0]
	url := ghAPIURL + "/repos/" + ghRepo + "/issues/" + prNumber + "/labels"
	fmt.Println("URL: ", url)

	timeout := time.Duration(5 * time.Second)
	client := &http.Client{
		Timeout: timeout,
	}

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	request.Header.Add("Accept", "application/vnd.github.v3+json")
	request.Header.Add("Authorization", "token "+ghToken)
	request.Header.Set("Content-Type", "application/json")
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := client.Do(request)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf(string(body))
	fmt.Println()
}

func main() {
	checkEnv()
	fmt.Println()

	var labelName string
	var labelColor string
	var labelDescription string

	var cmdLabel = &cobra.Command{
		Use:   "label [label]",
		Short: "Label a pull request",
		Long:  `label is for labeling a pull request.`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			postLabel(strings.Join(args, " "))
		},
	}

	cmdLabel.Flags().StringVarP(&labelName, "name", "n", "", "label name")
	cmdLabel.Flags().StringVarP(&labelColor, "color", "c", "", "label color")
	cmdLabel.Flags().StringVarP(&labelDescription, "description", "d", "", "label description")

	var rootCmd = &cobra.Command{Use: "app"}
	rootCmd.AddCommand(cmdLabel)
	rootCmd.Execute()
}

