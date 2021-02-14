package object

// Playlist represents PlaylistObject
// Link: https://developer.spotify.com/documentation/web-api/reference/#object-playlistobject
type Playlist struct {
	Name          string               `json:"name"`
	ID            string               `json:"id"`
	Collaborative bool                 `json:"collaborative"`
	Description   HTMLString           `json:"description"`
	ExternalURLs  ExternalURL          `json:"external_urls"`
	Followers     *Followers           `json:"followers,omitempty"`
	HRef          string               `json:"href"`
	Images        []Image              `json:"images,omitempty"`
	Owner         PublicUser           `json:"owner"`
	Public        bool                 `json:"public"`
	SnapshotID    SnapshotID           `json:"snapshot_id"`
	Tracks        PlaylistTracksCursor `json:"tracks"`
	Type          objectType           `json:"type"`
	URI           string               `json:"uri"`
}

// PlaylistTrack represents PlaylistTrackObject
// Link: https://developer.spotify.com/documentation/web-api/reference/#object-playlisttrackobject
type PlaylistTrack struct {
	AddedAt Timestamp   `json:"added_at"`
	AddedBy *PublicUser `json:"added_by"`
	IsLocal bool        `json:"is_local"`
	Track   Track       `json:"track"`
}

// PlaylistTracksCursor
type PlaylistTracksCursor struct {
	Paging
	Items []PlaylistTrack `json:"items"`
}
