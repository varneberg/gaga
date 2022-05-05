package main

import (
	// "context"
	"fmt"

	// "log"
	"os"
	// "syscall"
	// "golang.org/x/term"
)

var ghRef = os.Getenv("GITHUB_REF")
var ghRepo = os.Getenv("GITHUB_REPOSITORY")
var ghToken = os.Getenv("GITHUB_TOKEN")
var ghEventName = os.Getenv("GITHUB_EVENT_NAME")
var ghActor = os.Getenv("GITHUB_ACTOR")
var ghWorkflow = os.Getenv("GITHUB_WORKFLOW")
var ghActionsIDTokenRequestURL=os.Getenv("ACTIONS_ID_TOKEN_REQUEST_URL")
var ghActionsIDTokenRequestToken = os.Getenv("ACTIONS_ID_TOKEN_REQUEST_TOKEN")


func listEnv() {
	fmt.Println("Github Ref: ",ghRef)
	fmt.Println("Github Repo: ",ghRepo)
	fmt.Println("Github Token: ",ghToken)
	fmt.Println("Github Event: ",ghEventName)
	fmt.Println("Github Actor: ",ghActor)
	fmt.Println("Github Workflow: ",ghWorkflow)
	fmt.Println("Github Actions Token Request URL: ",ghActionsIDTokenRequestURL)
	fmt.Println("Github Actions Request Token: ",ghActionsIDTokenRequestToken)
	fmt.Printf("-----------------------\n")
}

func main() {
	listEnv()
	fmt.Println()
}
// func main() {
	//fmt.Print("GitHub Token: ")

	// for _, env := range os.Environ() {
	// 	fmt.Println(env)
	// }
	// byteToken, _ := term.ReadPassword(int(syscall.Stdin))
	// println()
	// token := string(byteToken)

	// ctx := context.Background()
	// ts := oauth2.StaticTokenSource(
	// 	&oauth2.Token{AccessToken: token},
	// )
	// tc := oauth2.NewClient(ctx, ts)

	// client := github.NewClient(tc)

	// user, resp, err := client.Users.Get(ctx, "")
	// if err != nil {
	// 	fmt.Printf("\nerror: %v\n", err)
	// 	return
	// }

	// // Rate.Limit should most likely be 5000 when authorized.
	// log.Printf("Rate: %#v\n", resp.Rate)

	// // If a Token Expiration has been set, it will be displayed.
	// if !resp.TokenExpiration.IsZero() {
	// 	log.Printf("Token Expiration: %v\n", resp.TokenExpiration)
	// }

	// fmt.Printf("\n%v\n", github.Stringify(user))
// }