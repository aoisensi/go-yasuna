package yasuna

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
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

func (t *Twitter) do(method, path string, uv url.Values, vo any) error {
	u := t.endpoint + path
	if uv != nil {
		u += "?" + uv.Encode()
	}
	req, _ := http.NewRequest(method, u, nil)
	req.Header.Set("Authorization", "Bearer "+t.Token.AccessToken)
	resp, err := t.cli.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	data, _ := io.ReadAll(resp.Body)
	//fmt.Println(string(data))
	code := resp.StatusCode
	if 200 > code || code >= 300 {
		return errors.New(string(data))
	}
	return json.Unmarshal(data, vo)
}

func (t *Twitter) get(path string, uv url.Values, vo any) error {
	return t.do("GET", path, uv, vo)
}

func (t *Twitter) delete(path string, uv url.Values, vo any) error {
	return t.do("DELETE", path, uv, vo)
}

func (t *Twitter) post(path string, vi, vo any) error {
	var buf *bytes.Buffer
	contentType := ""
	switch v := vi.(type) {
	case []byte:
		buf = bytes.NewBuffer(v)
	case url.Values:
		buf = bytes.NewBufferString(v.Encode())
		contentType = "application/x-www-form-urlencoded"
	default:
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
		return err
	}
	defer resp.Body.Close()
	data, _ := io.ReadAll(resp.Body)
	//fmt.Println(string(data))
	code := resp.StatusCode
	if 200 > code || code >= 300 {
		return errors.New(string(data))
	}
	return json.Unmarshal(data, vo)
}
