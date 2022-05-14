package yasuna

import "net/url"

type Expansion string

const (
	ExpansionAttachmentsPollIDs         Expansion = "attachments.poll_ids"
	ExpansionAttachmentsMediaKeys       Expansion = "attachments.media_keys"
	ExpansionAuthorID                   Expansion = "author_id"
	ExpansionEntitiesMentionsUsername   Expansion = "entities.mentions.username"
	ExpansionGeoPlaceID                 Expansion = "geo.place_id"
	ExpansionInReplyToUserID            Expansion = "in_reply_to_user_id"
	ExpansionReferencedTweetsID         Expansion = "referenced_tweets.id"
	ExpansionReferencedTweetsIDAuthorID Expansion = "referenced_tweets.id.author_id"
)

var ExpansionsAll = []Expansion{
	ExpansionAttachmentsPollIDs,
	ExpansionAttachmentsMediaKeys,
	ExpansionAuthorID,
	ExpansionEntitiesMentionsUsername,
	ExpansionGeoPlaceID,
	ExpansionInReplyToUserID,
	ExpansionReferencedTweetsID,
	ExpansionReferencedTweetsIDAuthorID,
}

type WithExpansions []Expansion

func (w WithExpansions) optGetTweet(v url.Values) {
	w.opt(v)
}

func (w WithExpansions) optGetUserTweets(v url.Values) {
	w.opt(v)
}

func (w WithExpansions) opt(v url.Values) {
	var s string
	for _, e := range w {
		s += string(e) + ","
	}
	v.Set("expansions", s[:len(s)-1])
}
