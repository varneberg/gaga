package comments

import (
	"encoding/json"
	// "fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/varneberg/gaga/requests"
)

var CommentCmd = &cobra.Command{
	Use:   "comment [comment]",
	Short: "Comment on pull request",
	Long:  `Commands related to commenting on pull requests`,
	Run: func(cmd *cobra.Command, args []string) {
		commentHandler()
	},
}

func parseComment(comment string) []byte {
	var body, err = json.Marshal(map[string]string{
		"body": comment,
	})
	if err != nil {
		log.Fatalln(err)
	}
	return body
}

func PostComment(comment string){
	body := parseComment(comment)
	requests.SendRequest("POST", requests.GetPrURL()+"/comments", body)

}

var comment string

func init() {
	CommentCmd.Flags().StringVarP(&comment, "new-comment", "n", "", "New comment on Pull Request")
}

func commentHandler(){
	if comment == ""{
		return
	}
	PostComment(comment)


}