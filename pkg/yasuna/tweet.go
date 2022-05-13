package yasuna

import (
	"net/http"
)

type Tweet struct {
	ID   int64  `json:"id,string"`
	Text string `json:"text"`
}

func (t *Twitter) PostTweet(params *PostTweetParams) (*Tweet, *http.Response, error) {
	out := new(struct {
		Data *Tweet `json:"data"`
	})
	resp, err := t.post("tweets/", params, out)
	return out.Data, resp, err
}

type IPostTweet interface {
	*PostTweetParams | string
}

type PostTweetParams struct {
	Text string `json:"text,omitempty"`
}
