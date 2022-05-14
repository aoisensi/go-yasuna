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
	DirectMessageDeepLink string                `json:"direct_message_deeplink,omitempty"`
	ForSuperFollowersOnly bool                  `json:"for_super_followers_only,omitempty"`
	Geo                   *PostTweetGeoParams   `json:"geo,omitempty"`
	Media                 *PostTweetMediaParams `json:"media,omitempty"`
	Poll                  *PostTweetPollParams  `json:"poll,omitempty"`
	QuoteTweetID          int64                 `json:"quote_tweet_id,string,omitempty"`
	Reply                 *PostTweetReplyParams `json:"reply,omitempty"`
	ReplySettings         string                `json:"reply_settings,omitempty"`
	Text                  string                `json:"text,omitempty"`
}

type PostTweetGeoParams struct {
	PlaceID string `json:"place_id,omitempty"`
}

type PostTweetMediaParams struct {
	MediaIDs      []string `json:"media_ids,omitempty"`
	TaggedUserIDs []int64  `json:"tagged_user_ids,string,omitempty"`
}

type PostTweetPollParams struct {
	DurationMinutes int      `json:"duration_minutes,omitempty"`
	Options         []string `json:"options,omitempty"`
}

type PostTweetReplyParams struct {
	ExcludeReplyUserIDs []int64 `json:"exclude_reply_user_ids,string,omitempty"`
	InReplyToTweetID    int64   `json:"in_reply_to_tweet_id,string,omitempty"`
}

// DeleteTweet allows a user or authenticated user ID to delete a Tweet.
//
// https://developer.twitter.com/en/docs/twitter-api/tweets/manage-tweets/api-reference/delete-tweets-id
func (t *Twitter) DeleteTweet(id int64) (*Response[*Deleted], error) {
	path := "tweets/" + strconv.FormatInt(id, 10)
	data := new(Response[*Deleted])
	return data, t.delete(path, nil, data)
}
