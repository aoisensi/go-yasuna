package yasuna

import "net/url"

// GetTweetsCountsRecent returns count of Tweets from the last seven days that match a query.
//
// https://developer.twitter.com/en/docs/twitter-api/tweets/counts/api-reference/get-tweets-counts-recent
func (t *Twitter) GetTweetsCountsRecent(query string, opts ...GetTweetsCountsOption) (*Response[[]*Counts], error) {
	return t.getTweetsCounts(query, "recent", opts...)
}

// This endpoint is only available to those users who have been approved for Academic Research access.
//
// The full-archive Tweet counts endpoint returns the count of Tweets that match your query from the complete history of public Tweets;
// since the first Tweet was created March 26, 2006.
//
// https://developer.twitter.com/en/docs/twitter-api/tweets/counts/api-reference/get-tweets-counts-all
func (t *Twitter) GetTweetsCountsAll(query string, opts ...GetTweetsCountsOption) (*Response[[]*Counts], error) {
	return t.getTweetsCounts(query, "all", opts...)
}

func (t *Twitter) getTweetsCounts(query, f string, opts ...GetTweetsCountsOption) (*Response[[]*Counts], error) {
	path := "tweets/counts/" + f
	uv := url.Values{
		"query": {query},
	}
	for _, opt := range opts {
		opt.optGetTweetsCounts(uv)
	}
	data := new(Response[[]*Counts])
	return data, t.get(path, uv, data)
}

type GetTweetsCountsOption interface {
	optGetTweetsCounts(v url.Values)
}
