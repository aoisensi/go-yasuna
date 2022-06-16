package yasuna

import "net/url"

type MediaField string

const (
	MediaFieldDurationMS       MediaField = "duration_ms"
	MediaFieldHeight           MediaField = "height"
	MediaFieldMediaKey         MediaField = "media_key"
	MediaFieldPreviewImageURL  MediaField = "preview_image_url"
	MediaFieldType             MediaField = "type"
	MediaFieldURL              MediaField = "url"
	MediaFieldWidth            MediaField = "width"
	MediaFieldPublicMetrics    MediaField = "public_metrics"
	MediaFieldNonPublicMetrics MediaField = "non_public_metrics"
	MediaFieldOrganicMetrics   MediaField = "organic_metrics"
	MediaFieldPromotedMetrics  MediaField = "promoted_metrics"
	MediaFieldAltText          MediaField = "alt_text"
)

var MediaFieldsAll = []MediaField{
	MediaFieldDurationMS,
	MediaFieldHeight,
	MediaFieldMediaKey,
	MediaFieldPreviewImageURL,
	MediaFieldType,
	MediaFieldURL,
	MediaFieldWidth,
	MediaFieldPublicMetrics,
	MediaFieldNonPublicMetrics,
	MediaFieldOrganicMetrics,
	MediaFieldPromotedMetrics,
	MediaFieldAltText,
}

type WithMediaFields []MediaField

func (w WithMediaFields) optGetTweet(v url.Values) {
	w.opt(v)
}

func (w WithMediaFields) optGetUserTweet(v url.Values) {
	w.opt(v)
}

func (w WithMediaFields) optGetTweetsSearch(v url.Values) {
	w.opt(v)
}

func (w WithMediaFields) opt(v url.Values) {
	var s string
	for _, e := range w {
		s += string(e) + ","
	}
	v.Set("media.fields", s[:len(s)-1])
}
