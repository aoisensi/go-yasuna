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
	tweets, err := twitter.GetUserTweets(id,
		yasuna.WithMaxResults(100),
		yasuna.WithExpansions([]yasuna.Expansion{yasuna.ExpansionAuthorID}),
		yasuna.WithTweetFields([]yasuna.TweetField{yasuna.TweetFieldAuthorID, yasuna.TweetFieldCreatedAt}),
		yasuna.WithUserFields([]yasuna.UserField{yasuna.UserFieldID}),
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, tweet := range tweets.Data {
		user := tweets.IncludeUser(tweet.AuthorID)
		fmt.Printf("@%s %s", user.Username, user.Name)
		fmt.Println(tweet.CreatedAt)
		fmt.Println(tweet.Text)
		fmt.Printf("http://twitter.com/%s/status/%d\n", user.Username, tweet.ID)
		fmt.Println("--------------------------------")
	}
}
