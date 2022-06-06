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
	}
	token := &yasuna.Token{
		RefreshToken: os.Getenv("TWITTER_REFRESH_TOKEN"),
		Scope:        yasuna.ScopeAll,
	}
	if err := oauth2.Refresh(token); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("AccessToken:", token.AccessToken)
	fmt.Println("RefreshToken:", token.RefreshToken)
}
