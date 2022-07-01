package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/varneberg/gaga/comments"
	"github.com/varneberg/gaga/tf"
	"github.com/varneberg/gaga/labels"
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
	
	var rootCmd = &cobra.Command{Use: "gaga"}
	rootCmd.AddCommand(labels.LabelCmd)
	rootCmd.AddCommand(tf.TFCmd)
	rootCmd.AddCommand(comments.CommentCmd)
	rootCmd.Execute()

}
