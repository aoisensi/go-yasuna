package yasuna

import "net/url"

type PlaceField string

const (
	PlaceFieldContainedWithin PlaceField = "contained_within"
	PlaceFieldCountry         PlaceField = "country"
	PlaceFieldFullName        PlaceField = "full_name"
	PlaceFieldGeo             PlaceField = "geo"
	PlaceFieldID              PlaceField = "id"
	PlaceFieldName            PlaceField = "name"
	PlaceFieldPlaceType       PlaceField = "place_type"
)

var PlaceFieldsAll = []PlaceField{
	PlaceFieldContainedWithin,
	PlaceFieldCountry,
	PlaceFieldFullName,
	PlaceFieldGeo,
	PlaceFieldID,
	PlaceFieldName,
	PlaceFieldPlaceType,
}

type WithPlaceFields []PlaceField

func (w WithPlaceFields) optGetTweet(v url.Values) {
	w.opt(v)
}

func (w WithPlaceFields) optGetUserTweets(v url.Values) {
	w.opt(v)
}

func (w WithPlaceFields) opt(v url.Values) {
	var s string
	for _, f := range w {
		s += string(f) + ","
	}
	v.Set("place.fields", s[:len(s)-1])
}
