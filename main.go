package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/varneberg/gaga/labels"
	"os"
)

// Check if any input was given
func checkArgs() {
	if len(os.Args[1:]) == 0 {
		fmt.Println("Error: No arguments")
		os.Exit(0)
	}
}

func init() {}

func main() {
	checkArgs()
	var rootCmd = &cobra.Command{Use: "app"}
	rootCmd.AddCommand(labels.LabelCmd)
	rootCmd.AddCommand(labels.TFCmd)
	rootCmd.Execute()

}
