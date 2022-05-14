package yasuna

import (
	"net/url"
	"strconv"
)

type WithMaxResults int

func (w WithMaxResults) optGetUserFollow(v url.Values) {
	v.Set("max_results", strconv.Itoa(int(w)))
}
