package yasuna

import (
	"net/url"
	"strconv"
	"time"
)

type WithEndTime time.Time

func (w WithEndTime) optGetUserTweet(v url.Values) {
	v.Set("end_time", time.Time(w).Format(time.RFC3339))
}

func (w WithEndTime) optGetTweetsSearch(v url.Values) {
	v.Set("end_time", time.Time(w).Format(time.RFC3339))
}

func (w WithEndTime) optGetTweetsCounts(v url.Values) {
	v.Set("end_time", time.Time(w).Format(time.RFC3339))
}

type WithMaxResults int

func (w WithMaxResults) optGetUserFollows(v url.Values) {
	v.Set("max_results", strconv.Itoa(int(w)))
}

func (w WithMaxResults) optGetUserTweet(v url.Values) {
	v.Set("max_results", strconv.Itoa(int(w)))
}

func (w WithMaxResults) optGetTweetsSearch(v url.Values) {
	v.Set("max_results", strconv.Itoa(int(w)))
}

func (w WithMaxResults) optGetUserFollow(v url.Values) {
	v.Set("max_results", strconv.Itoa(int(w)))
}

type WithNextToken string

func (w WithNextToken) optGetTweetsSearch(v url.Values) {
	v.Set("next_token", string(w))
}

func (w WithNextToken) optGetTweetsCounts(v url.Values) {
	v.Set("next_token", string(w))
}

type WithPaginationToken string

func (w WithPaginationToken) optGetUserTweet(v url.Values) {
	v.Set("pagination_token", string(w))
}

type WithSinceID int64

func (w WithSinceID) optGetUserTweet(v url.Values) {
	v.Set("since_id", strconv.FormatInt(int64(w), 10))
}

func (w WithSinceID) optGetTweetsSearch(v url.Values) {
	v.Set("since_id", strconv.FormatInt(int64(w), 10))
}

func (w WithSinceID) optGetTweetsCounts(v url.Values) {
	v.Set("since_id", strconv.FormatInt(int64(w), 10))
}

type WithStartTime time.Time

func (w WithStartTime) optGetUserTweet(v url.Values) {
	v.Set("start_time", time.Time(w).Format(time.RFC3339))
}

func (w WithStartTime) optGetTweetsSearch(v url.Values) {
	v.Set("start_time", time.Time(w).Format(time.RFC3339))
}

func (w WithStartTime) optGetTweetsCounts(v url.Values) {
	v.Set("start_time", time.Time(w).Format(time.RFC3339))
}

type WithUntilID int64

func (w WithUntilID) optGetUserTweet(v url.Values) {
	v.Set("until_id", strconv.FormatInt(int64(w), 10))
}

func (w WithUntilID) optGetTweetsSearch(v url.Values) {
	v.Set("until_id", strconv.FormatInt(int64(w), 10))
}

func (w WithUntilID) optGetTweetsCounts(v url.Values) {
	v.Set("until_id", strconv.FormatInt(int64(w), 10))
}
