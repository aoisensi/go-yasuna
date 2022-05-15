package yasuna

import "net/url"

type SortOrder string

const (
	SortOrderRecency   SortOrder = "recency"
	SortOrderRelevancy SortOrder = "relevancy"
)

type WithSortOrder SortOrder

func (w WithSortOrder) optGetTweetSearch(v url.Values) {
	v.Set("sort_order", string(w))
}
