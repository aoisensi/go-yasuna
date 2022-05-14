package yasuna

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type OAuth2 struct {
	ClientID     string
	ClientSecret string
	RedirectURI  string
	Client       *http.Client
	verifier     string
	state        string
}

func (o *OAuth2) AuthCodeURL(scope ...Scope) string {
	o.makeState()
	u := "https://twitter.com/i/oauth2/authorize?"
	u += url.Values{
		"response_type":         {"code"},
		"client_id":             {o.ClientID},
		"redirect_uri":          {o.RedirectURI},
		"scope":                 {string(fmtScope(scope...))},
		"state":                 {o.state},
		"code_challenge_method": {"S256"},
		"code_challenge":        {o.makePKCE()},
	}.Encode()
	return u
}

func (o *OAuth2) Exchange(code, state string) (*Twitter, error) {
	if state != o.state {
		return nil, errors.New("invalid state")
	}
	client := o.Client
	if client == nil {
		client = http.DefaultClient
	}

	body := strings.NewReader(url.Values{
		"code":          {code},
		"grant_type":    {"authorization_code"},
		"redirect_uri":  {o.RedirectURI},
		"code_verifier": {o.verifier},
	}.Encode())
	req, err := http.NewRequest("POST", "https://api.twitter.com/2/oauth2/token", body)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.SetBasicAuth(o.ClientID, o.ClientSecret)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	data, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(string(data))
	}
	token := new(Token)
	parseToken(data, token)
	return NewTwitter(client, token), nil
}

func (o *OAuth2) Refresh(token *Token) error {
	client := o.Client
	if client == nil {
		client = http.DefaultClient
	}
	body := strings.NewReader(url.Values{
		"grant_type":    {"refresh_token"},
		"refresh_token": {token.RefreshToken},
		"client_id":     {o.ClientID},
	}.Encode())
	req, err := http.NewRequest("POST", "https://api.twitter.com/2/oauth2/token", body)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.SetBasicAuth(o.ClientID, o.ClientSecret)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	data, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return errors.New(string(data))
	}
	parseToken(data, token)
	return nil
}

func (o *OAuth2) makePKCE() string {
	data := make([]byte, 48)
	rand.Read(data)
	o.verifier = base64.URLEncoding.EncodeToString(data)
	challenge := sha256.Sum256([]byte(o.verifier))
	c := base64.URLEncoding.WithPadding(0).EncodeToString(challenge[:])
	return c
}

func (o *OAuth2) makeState() {
	data := make([]byte, 6)
	rand.Read(data)
	o.state = base64.URLEncoding.EncodeToString(data)
}

type Token struct {
	Expire       *time.Time
	AccessToken  string
	RefreshToken string
	Scope        Scope
}

func parseToken(data []byte, token *Token) {
	v := struct {
		TokenType    string `json:"token_type"`
		ExpiresIn    int    `json:"expires_in"`
		AccessToken  string `json:"access_token"`
		Scope        string `json:"scope"`
		RefreshToken string `json:"refresh_token"`
	}{}
	if err := json.Unmarshal(data, &v); err != nil {
		panic(err)
	}
	token.AccessToken = v.AccessToken
	if v.ExpiresIn > 0 {
		expire := time.Now().Add(time.Duration(v.ExpiresIn) * time.Second)
		token.Expire = &expire
	} else {
		token.Expire = nil
	}
	if v.RefreshToken != "" {
		token.RefreshToken = v.RefreshToken
	}
	if v.Scope != "" {
		token.Scope = Scope(v.Scope)
	}
}
