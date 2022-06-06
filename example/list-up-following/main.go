package main

import (
	"fmt"
	"os"
	"strconv"

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
	id, _ := strconv.ParseInt(os.Getenv("TWITTER_USER_ID"), 10, 64)
	followers, err := twitter.GetUserFollowers(id, yasuna.WithMaxResults(1000))
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, follower := range followers.Data {
		fmt.Println(follower.Username)
	}
}
