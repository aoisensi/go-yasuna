package main

import (
	"fmt"
	"os"

	"github.com/aoisensi/go-yasuna/yasuna"
)

func main() {
	token := &yasuna.Token{
		AccessToken: os.Getenv("TWITTER_ACCESS_TOKEN"),
	}
	twitter := yasuna.NewTwitter(nil, nil, token)
	_, err := twitter.PostTweet(&yasuna.PostTweetParams{Text: "Hello, world!"})
	if err != nil {
		fmt.Println(err)
	}
}
