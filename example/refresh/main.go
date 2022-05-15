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
	}
	token := &yasuna.Token{
		AccessToken:  os.Getenv("TWITTER_ACCESS_TOKEN"),
		RefreshToken: os.Getenv("TWITTER_REFRESH_TOKEN"),
	}
	if err := oauth2.Refresh(token); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("AccessToken:", token.AccessToken)
	fmt.Println("RefreshToken:", token.RefreshToken)
}
