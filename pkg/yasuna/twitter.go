package yasuna

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type Twitter struct {
	endpoint string
	cli      *http.Client
	Token    *Token
}

func NewTwitter(client *http.Client, token *Token) *Twitter {
	if client == nil {
		client = http.DefaultClient
	}
	return &Twitter{
		endpoint: "https://api.twitter.com/2/",
		cli:      client,
		Token:    token,
	}
}

func (t *Twitter) post(path string, vi, vo any) (*http.Response, error) {
	var buf *bytes.Buffer
	contentType := ""
	if data, ok := vi.([]byte); ok {
		buf = bytes.NewBuffer(data)
	} else {
		data, err := json.Marshal(vi)
		if err != nil {
			panic(err)
		}
		buf = bytes.NewBuffer(data)
		contentType = "application/json"
	}
	req, _ := http.NewRequest("POST", t.endpoint+path, buf)
	req.Header.Set("Authorization", "Bearer "+t.Token.AccessToken)
	req.Header.Set("Content-Type", contentType)

	resp, err := t.cli.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	data, _ := io.ReadAll(resp.Body)
	code := resp.StatusCode
	if 200 > code || code >= 300 {
		return resp, errors.New(string(data))
	}
	return resp, json.Unmarshal(data, vo)
}
