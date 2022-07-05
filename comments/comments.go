package comments

import (
	"encoding/json"
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

func ToMarkdown(title string, body string) string{
	
	mdComment := "## " + title + "\n" + body
	return mdComment
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
var title string

func init() {
	CommentCmd.Flags().StringVarP(&comment, "new-comment", "n", "", "New comment on Pull Request")
	CommentCmd.Flags().StringVarP(&title, "title", "t", "", "Comment title (markdown)")
}

func commentHandler(){
	if comment == ""{
		return
	}
	mdComment := ToMarkdown(title, comment)
	PostComment(mdComment)
}