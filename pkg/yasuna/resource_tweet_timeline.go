package yasuna

import (
	"net/url"
	"strconv"
)

// GetUserTweets returns Tweets composed by a single user,
// specified by the requested user ID.
// By default, the most recent ten Tweets are returned per request.
// Using pagination, the most recent 3,200 Tweets can be retrieved.
//
// The Tweets returned by this endpoint count towards the Project-level Tweet cap.
//
// https://developer.twitter.com/en/docs/twitter-api/tweets/timelines/api-reference/get-users-id-tweets
func (t *Twitter) GetUserTweets(id int64, opt ...GetUserTweetsOption) (*Response[[]*Tweet], error) {
	return t.getUserTweets("tweets", id, opt...)
}

// GetUserMentions returns Tweets mentioning a single user specified by the requested user ID.
// By default, the most recent ten Tweets are returned per request.
// Using pagination, up to the most recent 800 Tweets can be retrieved.
//
// The Tweets returned by this endpoint count towards the Project-level Tweet cap.
//
// https://developer.twitter.com/en/docs/twitter-api/tweets/timelines/api-reference/get-users-id-mentions
func (t *Twitter) GetUserMentions(id int64, opt ...GetUserTweetsOption) (*Response[[]*Tweet], error) {
	return t.getUserTweets("mentions", id, opt...)
}

func (t *Twitter) getUserTweets(f string, id int64, opt ...GetUserTweetsOption) (*Response[[]*Tweet], error) {
	path := "users/" + strconv.FormatInt(id, 10) + "/" + f
	data := new(Response[[]*Tweet])
	vi := url.Values{}
	for _, o := range opt {
		o.optGetUserTweets(vi)
	}
	return data, t.get(path, vi, data)
}

type GetUserTweetsOption interface {
	optGetUserTweets(v url.Values)
}
