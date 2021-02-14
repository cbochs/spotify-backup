package auth

import (
	"context"
	"net/http"
	"os"

	"github.com/cbochs/spotify-backup-api/spotify"
	"github.com/pkg/errors"
	"golang.org/x/oauth2"
)

const (
	// AuthURL is
	AuthURL = "https://accounts.spotify.com/authorize"

	// TokenURL is
	TokenURL = "https://accounts.spotify.com/api/token"
)

type Config struct {
	config *oauth2.Config
}

func New() *Config {
	return &Config{
		config: &oauth2.Config{
			ClientID:     os.Getenv("SPOTIFY_CLIENTID"),
			ClientSecret: os.Getenv("SPOTIFY_CLIENTSECRET"),
			RedirectURL:  os.Getenv("SPOTIFY_REDIRECTURL"),
			Scopes:       []string{},
			Endpoint: oauth2.Endpoint{
				AuthURL:  AuthURL,
				TokenURL: TokenURL,
			},
		},
	}
}

func (c *Config) WithCredentials(clientID, clientSecret string) *Config {
	c.config.ClientID = clientID
	c.config.ClientSecret = clientSecret
	return c
}

func (c *Config) WithRedirect(redirectURL string) *Config {
	c.config.RedirectURL = redirectURL
	return c
}

func (c *Config) WithScopes(scopes ...string) *Config {
	if scopes != nil {
		c.config.Scopes = scopes
	}
	return c
}

func (c *Config) AuthCodeURL(state string, opts ...oauth2.AuthCodeOption) string {
	return c.config.AuthCodeURL(state, opts...)
}

func (c *Config) Exchange(ctx context.Context, code string, opts ...oauth2.AuthCodeOption) (*oauth2.Token, error) {
	return c.config.Exchange(ctx, code, opts...)
}

func (c *Config) Token(ctx context.Context, state string, req *http.Request, opts ...oauth2.AuthCodeOption) (*oauth2.Token, error) {
	values := req.URL.Query()

	if val := values.Get("error"); val != "" {
		return nil, errors.New("spotify: authorization returned an error: " + val)
	}
	if val := values.Get("state"); val != state {
		return nil, errors.New("spotify: authorization has mismatched state: " + val + ". Expected: " + state)
	}

	code := values.Get("code")
	if code == "" {
		return nil, errors.New("spotify: authorization returned empty code")
	}

	return c.Exchange(ctx, code, opts...)
}

func (c *Config) TokenSource(ctx context.Context, t *oauth2.Token) oauth2.TokenSource {
	return c.config.TokenSource(ctx, t)
}

func (c *Config) HTTPClient(ctx context.Context, t *oauth2.Token) *http.Client {
	return c.config.Client(ctx, t)
}

func (c *Config) Client(ctx context.Context, t *oauth2.Token) *spotify.Client {
	return spotify.New(c.HTTPClient(ctx, t))
}

func (c *Config) ClientWithSource(ctx context.Context, ts oauth2.TokenSource) *spotify.Client {
	return spotify.New(oauth2.NewClient(ctx, ts))
}
