package yasuna

import "strconv"

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
