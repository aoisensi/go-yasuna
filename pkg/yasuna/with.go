package yasuna

import (
	"net/url"
	"strconv"
	"time"
)

type WithEndTime time.Time

func (w WithEndTime) optGetUserTweets(v url.Values) {
	v.Set("end_time", time.Time(w).Format(time.RFC3339))
}

type WithMaxResults int

func (w WithMaxResults) optGetUserFollows(v url.Values) {
	v.Set("max_results", strconv.Itoa(int(w)))
}

func (w WithMaxResults) optGetUserTweets(v url.Values) {
	v.Set("max_results", strconv.Itoa(int(w)))
}

type WithPaginationToken string

func (w WithPaginationToken) optGetUserTweets(v url.Values) {
	v.Set("pagination_token", string(w))
}

type WithSinceID int64

func (w WithSinceID) optGetUserTweets(v url.Values) {
	v.Set("since_id", strconv.FormatInt(int64(w), 10))
}

type WithStartTime time.Time

func (w WithStartTime) optGetUserTweets(v url.Values) {
	v.Set("start_time", time.Time(w).Format(time.RFC3339))
}

type WithUntilID int64

func (w WithUntilID) optGetUserTweets(v url.Values) {
	v.Set("until_id", strconv.FormatInt(int64(w), 10))
}
