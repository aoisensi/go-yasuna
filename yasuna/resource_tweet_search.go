package yasuna

import "net/url"

// The recent search endpoint returns Tweets from the last seven days that match a search query.
//
// The Tweets returned by this endpoint count towards the Project-level Tweet cap.
//
// https://developer.twitter.com/en/docs/twitter-api/tweets/search/api-reference/get-tweets-search-recent
func (t *Twitter) GetTweetsSearchRecent(query string, opts ...GetTweetsSearchOption) (*Response[[]*Tweet], error) {
	return t.getTweetsSearch("recent", query, opts...)
}

// This endpoint is only available to those users who have been approved for Academic Research access.
//
// The full-archive search endpoint returns the complete history of public Tweets matching a search query; since the first Tweet was created March 26, 2006.
//
// The Tweets returned by this endpoint count towards the Project-level Tweet cap.
//
// https://developer.twitter.com/en/docs/twitter-api/tweets/search/api-reference/get-tweets-search-all
func (t *Twitter) GetTweetsSearchAll(query string, opts ...GetTweetsSearchOption) (*Response[[]*Tweet], error) {
	return t.getTweetsSearch("all", query, opts...)
}

func (t *Twitter) getTweetsSearch(f string, query string, opts ...GetTweetsSearchOption) (*Response[[]*Tweet], error) {
	path := "tweets/search/" + f
	data := new(Response[[]*Tweet])
	vi := url.Values{}
	vi.Set("query", query)
	for _, opt := range opts {
		opt.optGetTweetsSearch(vi)
	}
	return data, t.get(path, vi, data)
}

type GetTweetsSearchOption interface {
	optGetTweetsSearch(v url.Values)
}
