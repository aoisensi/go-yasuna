package main

import (
	"fmt"
	"os"

	"github.com/aoisensi/go-yasuna/yasuna"
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
