package labels

import "github.com/spf13/cobra"

var TFCmd = &cobra.Command{
	Use:   "tflabel",
	Short: "Create label from terraform plan",
	Long:  `Add labels based on output from terraform plan`,
	//Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		terraformParser()
	},
}

func init() {
	//LabelCmd.Flags().StringVarP(&labelName, "name", "n", "", "Label name")
	//LabelCmd.Flags().StringVarP(&labelColor, "color", "c", "", "Label color")
	//LabelCmd.Flags().StringVarP(&labelDescription, "description", "d", "", "Label description")
	//LabelCmd.PersistentFlags().BoolVarP(&labelRemove, "remove", "r", false, "Remove a label")
	//LabelCmd.PersistentFlags().BoolVarP(&removeAll, "remove-all", "R", false, "Remove all current labels on PR")
}

func terraformParser() {

}
