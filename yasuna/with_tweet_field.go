package yasuna

import "net/url"

type TweetField string

const (
	TweetFieldAttachments        TweetField = "attachments"
	TweetFieldAuthorID           TweetField = "author_id"
	TweetFieldContextAnnotations TweetField = "context_annotations"
	TweetFieldConversationID     TweetField = "conversation_id"
	TweetFieldCreatedAt          TweetField = "created_at"
	TweetFieldEntities           TweetField = "entities"
	TweetFieldGeo                TweetField = "geo"
	TweetFieldID                 TweetField = "id"
	TweetFieldInReplyToUserID    TweetField = "in_reply_to_user_id"
	TweetFieldLang               TweetField = "lang"
	TweetFieldNonPublicMetrics   TweetField = "non_public_metrics"
	TweetFieldPublicMetrics      TweetField = "public_metrics"
	TweetFieldOrganicMetrics     TweetField = "organic_metrics"
	TweetFieldPromotedMetrics    TweetField = "promoted_metrics"
	TweetFieldPossiblySensitive  TweetField = "possibly_sensitive"
	TweetFieldReferencedTweets   TweetField = "referenced_tweets"
	TweetFieldReplySettings      TweetField = "reply_settings"
	TweetFieldSource             TweetField = "source"
	TweetFieldText               TweetField = "text"
	TweetFieldWithheld           TweetField = "withheld"
)

var TweetFieldsAll = []TweetField{
	TweetFieldAttachments,
	TweetFieldAuthorID,
	TweetFieldContextAnnotations,
	TweetFieldConversationID,
	TweetFieldCreatedAt,
	TweetFieldEntities,
	TweetFieldGeo,
	TweetFieldID,
	TweetFieldInReplyToUserID,
	TweetFieldLang,
	TweetFieldNonPublicMetrics,
	TweetFieldPublicMetrics,
	TweetFieldOrganicMetrics,
	TweetFieldPromotedMetrics,
	TweetFieldPossiblySensitive,
	TweetFieldReferencedTweets,
	TweetFieldReplySettings,
	TweetFieldSource,
	TweetFieldText,
	TweetFieldWithheld,
}

type WithTweetFields []TweetField

func (w WithTweetFields) optGetTweet(v url.Values) {
	w.opt(v)
}

func (w WithTweetFields) optGetUserTweet(v url.Values) {
	w.opt(v)
}

func (w WithTweetFields) opt(v url.Values) {
	var s string
	for _, e := range w {
		s += string(e) + ","
	}
	v.Set("tweet.fields", s[:len(s)-1])
}
