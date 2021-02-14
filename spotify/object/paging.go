package object

// CursorPaging represents CursorPagingObject
// Link: https://developer.spotify.com/documentation/web-api/reference/#object-cursorpagingobject
type CursorPaging struct {
	Cursors Cursor `json:"cursors"`
	HRef    string `json:"href"`
	Limit   int    `json:"limit"`
	NextURL string `json:"next"`
	Total   int    `json:"total"`
}

// Cursor represents CursorObject
// Link: https://developer.spotify.com/documentation/web-api/reference/#object-cursorobject
type Cursor struct {
	After  string `json:"after,omitempty"`
	Before string `json:"before,omitempty"`
}

func (p *CursorPaging) Next() string {
	return p.NextURL
}

func (p *CursorPaging) Prev() string {
	return ""
}

// Paging represents PagingObject
// Link: https://developer.spotify.com/documentation/web-api/reference/#object-pagingobject
type Paging struct {
	HRef    string `json:"href"`
	Limit   int    `json:"limit"`
	NextURL string `json:"next"`
	Offset  int    `json:"offset"`
	PrevURL string `json:"previous"`
	Total   int    `json:"total"`
}

func (p *Paging) Next() string {
	return p.NextURL
}

func (p *Paging) Prev() string {
	return p.PrevURL
}
