package yasuna

import "net/url"

type PollField string

const (
	PollFieldDurationMinutes PollField = "duration_minutes"
	PollFieldEndDatetime     PollField = "end_datetime"
	PollFieldID              PollField = "id"
	PollFieldOptions         PollField = "options"
	PollFieldVotingStatus    PollField = "voting_status"
)

var PollFieldsAll = []PollField{
	PollFieldDurationMinutes,
	PollFieldEndDatetime,
	PollFieldID,
	PollFieldOptions,
	PollFieldVotingStatus,
}

type WithPollFields []PollField

func (w WithPollFields) optGetTweet(v url.Values) {
	w.opt(v)
}

func (w WithPollFields) optGetUserTweets(v url.Values) {
	w.opt(v)
}

func (w WithPollFields) opt(v url.Values) {
	var s string
	for _, f := range w {
		s += string(f) + ","
	}
	v.Set("poll.fields", s[:len(s)-1])
}
