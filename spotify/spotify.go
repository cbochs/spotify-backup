package spotify

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"

	"github.com/cbochs/spotify-backup-api/spotify/object"
	"github.com/cbochs/spotify-backup-api/spotify/options"
	"github.com/cbochs/spotify-backup-api/spotify/paging"
	"golang.org/x/oauth2"
)

// BaseURL is the base url of Spotify's API
// Link: https://developer.spotify.com/documentation/web-api/reference/
const BaseURL = "https://api.spotify.com/v1"

type Client struct {
	baseURL string
	client  *http.Client
	ts      oauth2.TokenSource
}

func New(c *http.Client) *Client {
	return &Client{
		baseURL: BaseURL,
		client:  c,
	}
}

// Me gets the current user's public profile
// Link: https://developer.spotify.com/documentation/web-api/reference/#endpoint-get-current-users-profile
func (c *Client) Me() (*object.PublicUser, error) {
	var user object.PublicUser
	if err := c.Get("/me", &user); err != nil {
		return nil, err
	}
	return &user, nil
}

func (c *Client) Playlist(id string) (*object.Playlist, error) {
	return c.PlaylistOpt(id, nil)
}

func (c *Client) PlaylistOpt(id string, opts *options.QueryOptions) (*object.Playlist, error) {
	var pl object.Playlist
	if err := c.GetOpt("/playlists"+id, &pl, opts); err != nil {
		return nil, err
	}
	return &pl, nil
}

func (c *Client) RecentlyPlayed() (*object.RecentlyPlayedCursor, error) {
	return c.RecentlyPlayedOpt(nil)
}

func (c *Client) RecentlyPlayedOpt(opts *options.QueryOptions) (*object.RecentlyPlayedCursor, error) {
	if opts != nil {
		opts.LimitMax(50)
	}

	url := c.EndpointURL("/me/player/recently-played", opts)

	var rp object.RecentlyPlayedCursor
	if err := c.GetOpt(url, &rp, opts); err != nil {
		return nil, err
	}
	return &rp, nil
}

func (c *Client) SavedTracks() (*object.SavedTracksCursor, error) {
	return c.SavedTracksOpt(nil)
}

func (c *Client) SavedTracksOpt(opts *options.QueryOptions) (*object.SavedTracksCursor, error) {
	if opts != nil {
		opts.LimitMax(50)
	}

	url := c.EndpointURL("/me/tracks", opts)

	var st object.SavedTracksCursor
	if err := c.GetOpt(url, &st, opts); err != nil {
		return nil, err
	}
	return &st, nil
}

func (c *Client) Next(page paging.Paging) error {
	next := page.Next()
	if next == "" {
		return NoMorePagesError
	}

	v := reflect.ValueOf(page).Elem()
	v.Set(reflect.Zero(v.Type()))

	return c.Get(next, page)
}

func (c *Client) Prev(page paging.Paging) error {
	prev := page.Prev()
	if prev == "" {
		return NoMorePagesError
	}

	v := reflect.ValueOf(page).Elem()
	v.Set(reflect.Zero(v.Type()))

	return c.Get(prev, page)
}

func (c *Client) Get(url string, v interface{}) error {
	return c.GetOpt(url, v, nil)
}

func (c *Client) GetOpt(url string, v interface{}, opts *options.QueryOptions) error {
	fmt.Printf("Getting: %s\n", url)
	resp, err := c.client.Get(url)
	if err != nil {
		return &Error{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	defer resp.Body.Close()

	byt, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return &Error{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	if resp.StatusCode > 299 {
		var err object.Error
		if err := json.Unmarshal(byt, &err); err != nil {
			return &Error{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
			}
		}

		return &Error{
			Message: err.Error.Message,
			Status:  err.Error.Status,
		}
	}

	if err := json.Unmarshal(byt, v); err != nil {
		return &Error{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return nil
}

func (c *Client) EndpointURL(endpoint string, opts *options.QueryOptions) string {
	query := ""
	if opts != nil {
		query = opts.Values().Encode()
	}
	if query != "" {
		return fmt.Sprintf("%s%s?%s", c.baseURL, endpoint, query)
	}
	return fmt.Sprintf("%s%s", c.baseURL, endpoint)
}
