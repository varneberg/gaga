package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/varneberg/gaga/labels"
	//"github.com/varneberg/gaga/labels"
	"os"
)

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
