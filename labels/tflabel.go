package labels

import "github.com/spf13/cobra"

var TFCmd = &cobra.Command{
	Use:   "tflabel",
	Short: "Labels from terraform plan",
	Long:  `Add labels based on terraform plan from unix pipe.`,
	//Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		terraformParser()
	},
}

var labelAddUpdate string
var labelDestroy string
var labelNoChanges string
var labelError string

func init() {
	TFCmd.Flags().StringVarP(&labelAddUpdate, "label-add", "a", "tf/add-update", "Terraform Add/Update label name")
	TFCmd.Flags().StringVarP(&labelDestroy, "label-destroy", "d", "tf/destroy", "Terraform destroy label name")
	TFCmd.Flags().StringVarP(&labelNoChanges, "label-no-changes", "n", "tf/no-changes", "Terraform no changes label name")
	TFCmd.Flags().StringVarP(&labelError, "label-error", "e", "tf/error", "Terraform error label name")
	//LabelCmd.Flags().StringVarP(&labelName, "name", "n", "", "Label name")
	//LabelCmd.Flags().StringVarP(&labelColor, "color", "c", "", "Label color")
	//LabelCmd.Flags().StringVarP(&labelDescription, "description", "d", "", "Label description")
	//LabelCmd.PersistentFlags().BoolVarP(&labelRemove, "remove", "r", false, "Remove a label")
	//LabelCmd.PersistentFlags().BoolVarP(&removeAll, "remove-all", "R", false, "Remove all current labels on PR")
}

func terraformParser() {

}
