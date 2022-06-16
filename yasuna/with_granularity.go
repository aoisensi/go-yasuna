package yasuna

import "net/url"

type Granularity string

const (
	GranularityMinute Granularity = "minute"
	GranularityHour   Granularity = "hour"
	GranularityDay    Granularity = "day"
)

type WithGranularity Granularity

func (w WithGranularity) optGetTweetsCounts(v url.Values) {
	v.Set("granularity", string(w))
}
