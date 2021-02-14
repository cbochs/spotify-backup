package object

// SavedTrack represents SavedTrackObject
// Link: https://developer.spotify.com/documentation/web-api/reference/#object-savedtrackobject
type SavedTrack struct {
	AddedAt Timestamp `json:"added_at"`
	Track   Track     `json:"track"`
}

// SavedTracksCursor
type SavedTracksCursor struct {
	Paging
	Items []SavedTrack `json:"items"`
}
