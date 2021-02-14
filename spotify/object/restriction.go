package object

// AlbumRestriction represents AlbumRestrictionObject
// Link: https://developer.spotify.com/documentation/web-api/reference/#object-albumrestrictionobject
type AlbumRestriction struct {
	Reason string `json:"reason"`
}

// TrackRestriction represents TrackRestrictionObject
// Link: https://developer.spotify.com/documentation/web-api/reference/#object-trackrestrictionobject
type TrackRestriction struct {
	Reason string `json:"reason"`
}

// RestrictionType represents the different reasons for a tracks restriction
type RestrictionType string

//
const (
	MarketRestriction   RestrictionType = "market"
	ProductRestriction  RestrictionType = "product"
	ExplicitRestriction RestrictionType = "explicit"
	// UnknownRestriction RestrictionType
)
