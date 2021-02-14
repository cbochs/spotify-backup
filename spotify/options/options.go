package options

import (
	"math"
	"net/url"
	"strconv"
	"strings"
)

type QueryOptions struct {
	IDs    []string
	maxIDs int

	Market    string
	Country   string
	Locale    string
	Timestamp string

	// Paging
	limit    int
	limitMax int
	Offset   int
	After    int
	Before   int

	// Recommendations
	SeedArtists      []string
	SeedGenres       []string
	SeedTracks       []string
	Acousticness     *featureBounds
	Danceability     *featureBounds
	DurationMS       *featureBounds
	Energy           *featureBounds
	Instrumentalness *featureBounds
	Key              *featureBounds
	Liveness         *featureBounds
	Loudness         *featureBounds
	Mode             *featureBounds
	Popularity       *featureBounds
	Speechiness      *featureBounds
	Tempo            *featureBounds
	TimeSignature    *featureBounds
	Valence          *featureBounds

	// Personalization API
	TimeRange string

	// Response Customization
	fields          string
	AdditionalTypes []string

	// Player API
	URI           string
	State         *bool
	PositionMS    int
	VolumePercent int
	DeviceID      string

	// Playlists API
	Position int
	URIs     []string

	// Search API
	Q               string
	Type            string
	IncludeExternal string
}

type featureBounds struct {
	Min    int
	Max    int
	Target int
}

func Query() *QueryOptions {
	return &QueryOptions{}
}

func (q *QueryOptions) Limit(n int) *QueryOptions {
	q.limit = n
	return q
}

func (q *QueryOptions) LimitMax(max int) *QueryOptions {
	q.limitMax = max
	return q
}

func (q *QueryOptions) Fields(fields string) *QueryOptions {
	q.fields = fields
	return q
}

func (q *QueryOptions) Values() *url.Values {
	v := &url.Values{}

	if q.IDs != nil {
		max := len(q.IDs)
		if q.maxIDs > 0 {
			max = clamp(0, q.maxIDs, max)
		}
		v.Add("ids", strings.Join(q.IDs[:max], ","))
	}

	if q.Market != "" {
		v.Add("market", q.Market)
	}
	if q.Country != "" {
		v.Add("country", q.Country)
	}
	if q.Locale != "" {
		v.Add("locale", q.Locale)
	}
	if q.Timestamp != "" {
		v.Add("timestamp", q.Timestamp)
	}

	if q.limit > 0 {
		max := q.limit
		if q.limitMax > 0 {
			max = q.limitMax
		}
		v.Add("limit", strconv.Itoa(clamp(0, max, q.limit)))
	}
	if q.Offset > 0 {
		v.Add("offset", strconv.Itoa(q.Offset))
	}
	if q.After > 0 {
		v.Add("after", strconv.Itoa(q.After))
	}
	if q.Before > 0 {
		v.Add("before", strconv.Itoa(q.Before))
	}

	if q.fields != "" {
		v.Add("fields", q.fields)
	}

	return v
}

func clamp(min, max, value int) int {
	clampMin := math.Max(float64(min), float64(value))
	clampMax := math.Min(float64(max), clampMin)
	return int(clampMax)
}
