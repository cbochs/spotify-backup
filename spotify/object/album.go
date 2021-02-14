package object

import "time"

// SimplifiedAlbum represents SimplifiedAlbumObject
// Link: https://developer.spotify.com/documentation/web-api/reference/#object-simplifiedalbumobject
type SimplifiedAlbum struct {
	Name                 string             `json:"name"`
	ID                   string             `json:"id"`
	AlbumGroup           string             `json:"album_group,omitempty"`
	AlbumType            string             `json:"album_type"`
	Artists              []SimplifiedArtist `json:"artists"`
	AvailableMarkets     []string           `json:"available_markets"`
	ExternalURLs         ExternalURL        `json:"external_urls"`
	HRef                 string             `json:"href"`
	Images               []Image            `json:"images"`
	ReleaseDate          string             `json:"release_date"` // TODO: use ReleaseDate struct
	ReleaseDatePrecision string             `json:"release_date_precision"`
	Restrictions         *AlbumRestriction  `json:"restrictions,omitempty"`
	Type                 string             `json:"type"`
	URI                  string             `json:"uri"`
}

// ReleaseDate is a coarse timestamp in either YYYY, YYYY-MM, or YYYY-MM-DD format
type ReleaseDate struct {
	time.Time
}

// MarshalJSON implements serialization for a release date
// func (rd ReleaseDate) MarshalJSON() ([]byte, error) {}

// UnmarshalJSON implements deserialization for a release datee
// func (rd *ReleaseDate) UnmarshalJSON(b []byte) error {}
