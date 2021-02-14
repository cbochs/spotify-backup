package object

// Context represents ContextObject
// Link: https://developer.spotify.com/documentation/web-api/reference/#object-contextobject
type Context struct {
	ExternalURLs ExternalURL `json:"external_urls"`
	HRef         string      `json:"href"`
	Type         objectType  `json:"type"`
	URI          string      `json:"uri"`
}

// ExternalID represents ExternalIdObject
// Link: https://developer.spotify.com/documentation/web-api/reference/#object-externalidobject
type ExternalID struct {
	EAN  string `json:"ean,omitempty"`  // International Articla Number
	ISRC string `json:"isrc,omitempty"` // International Standard Recording Code
	UPC  string `json:"upc,omitempty"`  // Universal Product Code
}

// ExternalURL represents ExternalUrlObject
// Link: https://developer.spotify.com/documentation/web-api/reference/#object-externalurlobject
type ExternalURL struct {
	Spotify string `json:"spotify"`
}

// Followers represents FollowersObject
// Link: https://developer.spotify.com/documentation/web-api/reference/#object-followersobject
type Followers struct {
	HRef  string `json:"href"`
	Total int    `json:"total"`
}

// Image represents ImageObject
// :ink: https://developer.spotify.com/documentation/web-api/reference/#object-imageobject
type Image struct {
	URL    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

// ExplicitContentSettings represents ExplicitContentSettingsObject
// Link: https://developer.spotify.com/documentation/web-api/reference/#object-explicitcontentsettingsobject
type ExplicitContentSettings struct {
	FilterEnabled bool `json:"filter_enabled"`
	FilterLocked  bool `json:"filter_locked"`
}

// objectType represents the different objects that are available from Spotify's API
// Permitted values: "user", "track", "playlist", ...
type objectType string

// These are the object types defined by the Spotify API
// AlbumType:    A
// ArtistType:   A
// PlaylistType: SimplifiedPlaylist, Playlist
// ShowType:     A
// TrackType:    SimplifiedTrack, Track, and LinkedTrack
// UserType:     PrivateUser, PublicUser
const (
	AlbumType    objectType = "album"
	ArtistType   objectType = "artist"
	PlaylistType objectType = "playlist"
	ShowType     objectType = "show"
	TrackType    objectType = "track"
	UserType     objectType = "user"
)
