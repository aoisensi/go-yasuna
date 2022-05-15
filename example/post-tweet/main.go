package main

import (
	"fmt"
	"os"

	"github.com/aoisensi/go-yasuna/pkg/yasuna"
)

func main() {
	oauth2 := &yasuna.OAuth2{
		ClientID:     os.Getenv("TWITTER_CLIENT_ID"),
		ClientSecret: os.Getenv("TWITTER_CLIENT_SECRET"),
		RedirectURI:  os.Getenv("TWITTER_REDIRECT_URL"),
	}
	token := &yasuna.Token{
		AccessToken: os.Getenv("TWITTER_ACCESS_TOKEN"),
	}
	twitter := yasuna.NewTwitter(nil, oauth2, token)
	_, err := twitter.PostTweet(&yasuna.PostTweetParams{Text: "Hello, world!"})
	if err != nil {
		fmt.Println(err)
	}
}
