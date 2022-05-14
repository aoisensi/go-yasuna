package main

import (
	"fmt"
	"os"

	"github.com/aoisensi/go-yasuna/pkg/yasuna"
)

func main() {
	token := &yasuna.Token{
		AccessToken: os.Getenv("TWITTER_ACCESS_TOKEN"),
	}
	twitter := yasuna.NewTwitter(nil, token)
	tweet, err := twitter.GetTweet(20)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(tweet.Data.Text)
}
