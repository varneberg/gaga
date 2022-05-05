package main

import (
	// "context"
	"fmt"
	// "log"
	"os"
	// "syscall"

	// "github.com/google/go-github/v44/github"
	// "golang.org/x/oauth2"
	// "golang.org/x/term"
)

func main() {
	//fmt.Print("GitHub Token: ")

	for _, env := range os.Environ() {
		fmt.Println(env)
	}
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
}