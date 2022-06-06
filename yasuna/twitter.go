package yasuna

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

type Twitter struct {
	endpoint string
	cli      *http.Client
	Token    *Token
	OAuth2   *OAuth2
}

func NewTwitter(client *http.Client, oauth2 *OAuth2, token *Token) *Twitter {
	if client == nil {
		client = http.DefaultClient
	}
	return &Twitter{
		endpoint: "https://api.twitter.com/2/",
		cli:      client,
		Token:    token,
		OAuth2:   oauth2,
	}
}

func (t *Twitter) get(path string, uv url.Values, vo any) error {
	return t.do("GET", path, uv, vo)
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

	if err := t.refresh(); err != nil {
		return err
	}

	resp, err := t.cli.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return t.done(resp, vo)
}

func (t *Twitter) delete(path string, uv url.Values, vo any) error {
	return t.do("DELETE", path, uv, vo)
}

func (t *Twitter) do(method, path string, uv url.Values, vo any) error {
	u := t.endpoint + path
	if uv != nil {
		u += "?" + uv.Encode()
	}
	req, _ := http.NewRequest(method, u, nil)
	req.Header.Set("Authorization", "Bearer "+t.Token.AccessToken)

	if err := t.refresh(); err != nil {
		return err
	}

	resp, err := t.cli.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return t.done(resp, vo)
}

func (t *Twitter) done(resp *http.Response, vo any) error {
	code := resp.StatusCode
	data, _ := io.ReadAll(resp.Body)
	//fmt.Println(string(data))
	if 200 > code || code >= 300 {
		err := new(ResponseError)
		if err := json.Unmarshal(data, err); err != nil {
			return err
		}
		return err
	}
	return json.Unmarshal(data, vo)
}

func (t *Twitter) refresh() error {
	if t.OAuth2 == nil {
		return nil
	}
	if t.Token.RefreshToken == "" {
		return nil
	}
	if t.Token.Expired() {
		return t.OAuth2.Refresh(t.Token)
	}
	return nil
}
