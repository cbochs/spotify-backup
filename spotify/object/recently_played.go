package object

// PlayHistory represents PlayHistoryObject
// Link: https://developer.spotify.com/documentation/web-api/reference/#object-playhistoryobject
type PlayHistory struct {
	Context  *Context        `json:"context"`
	PlayedAt Timestamp       `json:"played_at"`
	Track    SimplifiedTrack `json:"track"`
}

// RecentlyPlayedCursor
type RecentlyPlayedCursor struct {
	CursorPaging
	Items []PlayHistory `json:"items"`
}
