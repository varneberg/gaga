package main

import (
	// "context"
	"errors"
	"fmt"
	"strings"

	// "log"
	"os"
	// "syscall"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	// "golang.org/x/term"
)

const EnvBaseURL = "GITHUB_BASE_URL"
type Config struct {
	Token        string
	BaseURL      string
	Owner        string
	Repo         string
	CI           string
}

type Client struct {
	*github.Client
	Debug bool

	Config Config

	common service
}

type service struct {
	client *Client
}

func NewClient() (*Client, error) {
	var cfg Config
	token := cfg.Token

	if strings.HasPrefix(token, "$") {
		token = os.Getenv(strings.TrimPrefix(token, "$"))
	}

	if token == "" {
		fmt.Println("github token is missing")
		return &Client{}, errors.New("github token is missing")
	}

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)
	client := github.NewClient(tc)

	baseURL := cfg.BaseURL
	baseURL = strings.TrimPrefix(baseURL, "$")
	if baseURL == EnvBaseURL {
		baseURL = os.Getenv(EnvBaseURL)
	}
	if baseURL != "" {
		var err error
		client, err = github.NewEnterpriseClient(baseURL, baseURL, tc)
		if err != nil {
			fmt.Println("failed to create a new github api client")
			return &Client{}, errors.New("failed to create a new github api client")
		}
	} 
	c := &Client{
		Config: cfg,
		Client: client,
	}
	return c, nil
}

func listEnv() {
	for _, env := range os.Environ() {
		fmt.Println(env)
	}
	fmt.Printf("-----------------------\n")
}

func main() {
	listEnv()
	fmt.Println()
	NewClient()
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