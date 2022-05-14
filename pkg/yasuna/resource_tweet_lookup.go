package yasuna

import (
	"net/url"
	"strconv"
)

// GetTweet returns a variety of information about a single Tweet specified by the requested ID.
//
// https://developer.twitter.com/en/docs/twitter-api/tweets/lookup/api-reference/get-tweets-id
func (t *Twitter) GetTweet(id int64, opt ...GetTweetOption) (*Response[*Tweet], error) {
	path := "tweets/" + strconv.FormatInt(id, 10)
	uv := url.Values{}
	for _, o := range opt {
		o.optGetTweet(uv)
	}
	data := new(Response[*Tweet])
	return data, t.get(path, uv, data)
}

// GetTweets returns a variety of information about the Tweet specified by the requested ID or list of IDs.
//
// https://developer.twitter.com/en/docs/twitter-api/tweets/lookup/api-reference/get-tweets
func (t *Twitter) GetTweets(ids []int64, opt ...GetTweetOption) (*Response[[]*Tweet], error) {
	path := "tweets"
	uv := url.Values{"id": {fmtIds(ids)}}
	for _, o := range opt {
		o.optGetTweet(uv)
	}
	data := new(Response[[]*Tweet])
	return data, t.get(path, uv, data)
}

type GetTweetOption interface {
	optGetTweet(v url.Values)
}
