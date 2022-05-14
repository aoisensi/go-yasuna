package yasuna

type ReferencedTweet struct {
	Type string `json:"type"`
	ID   int64  `json:"id,string"`
}
