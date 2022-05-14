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
	deleted, err := twitter.DeleteTweet(1525477480705888256)
	if err != nil {
		fmt.Println(err)
		return
	}
	if deleted.Data.Deleted {
		fmt.Println("Tweet deleted.")
	} else {
		fmt.Println("Tweet not deleted.")
	}
}
