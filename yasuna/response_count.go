package yasuna

import "time"

type Counts struct {
	Start      time.Time `json:"start"`
	End        time.Time `json:"end"`
	TweetCount int       `json:"tweet_count"`
}
