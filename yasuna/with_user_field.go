package yasuna

import "net/url"

type UserField string

const (
	UserFieldCreatedAt       UserField = "created_at"
	UserFieldDescription     UserField = "description"
	UserFieldEntities        UserField = "entities"
	UserFieldID              UserField = "id"
	UserFieldLocation        UserField = "location"
	UserFieldPinnedTweetID   UserField = "pinned_tweet_id"
	UserFieldProfileImageURL UserField = "profile_image_url"
	UserFieldProtected       UserField = "protected"
	UserFieldPublicMetrics   UserField = "public_metrics"
	UserFieldURL             UserField = "url"
	UserFieldUsername        UserField = "username"
	UserFieldVerified        UserField = "verified"
	UserFieldWithheld        UserField = "withheld"
)

var UserFieldsAll = []UserField{
	UserFieldCreatedAt,
	UserFieldDescription,
	UserFieldEntities,
	UserFieldID,
	UserFieldLocation,
	UserFieldPinnedTweetID,
	UserFieldProfileImageURL,
	UserFieldProtected,
	UserFieldPublicMetrics,
	UserFieldURL,
	UserFieldUsername,
	UserFieldVerified,
	UserFieldWithheld,
}

type WithUserFields []UserField

func (w WithUserFields) optGetTweet(v url.Values) {
	w.opt(v)
}

func (w WithUserFields) optGetUserTweet(v url.Values) {
	w.opt(v)
}

func (w WithUserFields) optGetTweetsSearch(v url.Values) {
	w.opt(v)
}

func (w WithUserFields) opt(v url.Values) {
	var s string
	for _, e := range w {
		s += string(e) + ","
	}
	v.Set("user.fields", s[:len(s)-1])
}
