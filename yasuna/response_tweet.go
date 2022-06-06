package yasuna

import "time"

type Tweet struct {
	ID                 int64              `json:"id,string"`
	Text               string             `json:"text"`
	Attachment         *Attachment        `json:"attachment,omitempty"`
	AuthorID           int64              `json:"author_id,string"`
	ContextAnnotations []any              `json:"context_annotations,omitempty"`
	ConversationID     int64              `json:"conversation_id,string"`
	CreatedAt          time.Time          `json:"created_at"`
	Entities           *Entities          `json:"entities,omitempty"`
	Geo                *Geo               `json:"geo,omitempty"`
	InReplyToUserID    int64              `json:"in_reply_to_user_id,string"`
	Lang               string             `json:"lang,omitempty"`
	OrganicMetrics     *OrganicMetrics    `json:"organic_metrics,omitempty"`
	PositiveSensitive  bool               `json:"positive_sensitive"`
	PromotedMetrics    *PromotedMetrics   `json:"promoted_metrics,omitempty"`
	ReferencedTweet    []*ReferencedTweet `json:"referenced_tweet,omitempty"`
	ReplySettings      string             `json:"reply_settings"`
	Source             string             `json:"source"`
	Withheld           *Withheld          `json:"withheld,omitempty"`
}
