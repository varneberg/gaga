package main

import (
	"fmt"
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

func main() {
	checkArgs()
	switch os.Args[1] {
	case "label":
		labels.LabelHandler(os.Args[2:])
	}
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
	//	newLabel := newLabel{
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
