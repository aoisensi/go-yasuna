package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/aoisensi/go-yasuna/pkg/yasuna"
)

func main() {
	token := &yasuna.Token{
		AccessToken: os.Getenv("TWITTER_ACCESS_TOKEN"),
	}
	twitter := yasuna.NewTwitter(nil, token)
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
