package object

// SimplifiedArtist represents SimplifiedArtistObject
// Link: https://developer.spotify.com/documentation/web-api/reference/#object-simplifiedartistobject
type SimplifiedArtist struct {
	Name         string      `json:"name"`
	ID           string      `json:"id"`
	ExternalURLs ExternalURL `json:"external_urls"`
	HRef         string      `json:"href"`
	Type         string      `json:"type"`
	URI          string      `json:"uri"`
}

// Artist represents ArtistObject
// Link: https://developer.spotify.com/documentation/web-api/reference/#object-artistobject
type Artist struct {
	SimplifiedArtist
	Followers  *Followers `json:"followers,omitempty"`
	Genres     []string   `json:"genres,omitempty"`
	Images     []Image    `json:"images,omitempty"`
	Popularity int        `json:"popularity"`
}
