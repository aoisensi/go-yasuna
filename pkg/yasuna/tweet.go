package yasuna

import (
	"net/url"
	"strconv"
)

// GetTweet returns a variety of information about a single Tweet specified by the requested ID.
//
// https://developer.twitter.com/en/docs/twitter-api/tweets/lookup/api-reference/get-tweets-id
func (t *Twitter) GetTweet(id int64) (*Response[*Tweet], error) {
	path := "tweets/" + strconv.FormatInt(id, 10)
	uv := url.Values{
		"expansions":   {"author_id"},
		"tweet.fields": {"id,text,author_id,created_at"},
		"user.fields":  {"username,verified"},
	}
	data := new(Response[*Tweet])
	return data, t.get(path, uv, data)
}

// GetTweets returns a variety of information about the Tweet specified by the requested ID or list of IDs.
//
// https://developer.twitter.com/en/docs/twitter-api/tweets/lookup/api-reference/get-tweets
func (t *Twitter) GetTweets(ids []int64) (*Response[[]*Tweet], error) {
	path := "tweets"
	uv := url.Values{
		"id":           {convIds(ids)},
		"expansions":   {"author_id"},
		"tweet.fields": {"id,text,author_id,created_at"},
		"user.fields":  {"username,verified"},
	}
	data := new(Response[[]*Tweet])
	return data, t.get(path, uv, data)
}

// PostTweet creates a Tweet on behalf of an authenticated user.
//
// https://developer.twitter.com/en/docs/twitter-api/tweets/manage-tweets/api-reference/post-tweets
func (t *Twitter) PostTweet(params *PostTweetParams) (*Response[*Tweet], error) {
	data := new(Response[*Tweet])
	return data, t.post("tweets/", params, data)
}

type PostTweetParams struct {
	Text string `json:"text,omitempty"`
}

// DeleteTweet allows a user or authenticated user ID to delete a Tweet.
//
// https://developer.twitter.com/en/docs/twitter-api/tweets/manage-tweets/api-reference/delete-tweets-id
func (t *Twitter) DeleteTweet(id int64) (*Response[*Deleted], error) {
	path := "tweets/" + strconv.FormatInt(id, 10)
	data := new(Response[*Deleted])
	return data, t.delete(path, nil, data)
}
