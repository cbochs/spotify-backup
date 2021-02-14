package object

// LinkedTrack represents LinkedTrackObject
// Link: https://developer.spotify.com/documentation/web-api/reference/#object-linkedtrackobject
type LinkedTrack struct {
	ID           string      `json:"id"`
	ExternalURLs ExternalURL `json:"external_urls"`
	HRef         string      `json:"href"`
	Type         string      `json:"type"`
	URI          string      `json:"uri"`
}

// SimplifiedTrack represents SimplifiedTrackObject
// Link: https://developer.spotify.com/documentation/web-api/reference/#object-simplifiedtrackobject
type SimplifiedTrack struct {
	Name             string             `json:"name"`
	ID               string             `json:"id"`
	Artists          []SimplifiedArtist `json:"artists"`
	AvailableMarkets []string           `json:"available_markets"`
	DiscNumber       int                `json:"disc_number"`
	DurationMS       int                `json:"duration_ms"`
	Explicit         bool               `json:"explicit"`
	ExternalURLs     ExternalURL        `json:"external_urls"`
	HRef             string             `json:"href"`
	IsLocal          bool               `json:"is_local"`
	IsPlayable       *bool              `json:"is_playable,omitempty"`
	LinkedFrom       *LinkedTrack       `json:"linked_from,omitempty"`
	PreviewURL       string             `json:"preview_url"`
	Restrictions     *TrackRestriction  `json:"restrictions,omitempty"`
	TrackNumber      int                `json:"track_number"`
	Type             string             `json:"type"`
	URI              string             `json:"uri"`
}

// Track represents TrackObject
// Link: https://developer.spotify.com/documentation/web-api/reference/#object-trackobject
type Track struct {
	SimplifiedTrack
	Album       SimplifiedAlbum    `json:"album"`
	Artists     []SimplifiedArtist `json:"artists"`
	ExternalIDs ExternalID         `json:"external_ids"`
	Popularity  int                `json:"popularity"`
}
