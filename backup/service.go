package backup

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/cbochs/spotify-backup-api/backupdb"
	"github.com/cbochs/spotify-backup-api/schema"
	"github.com/cbochs/spotify-backup-api/spotify/auth"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/oauth2"
)

type Service struct {
	ctx             context.Context
	db              *backupdb.Client
	spa             *auth.Config
	clients         map[string]*Client
	mu              sync.Mutex
	omitDisplayName bool
}

func NewService(ctx context.Context, db *backupdb.Client, spa *auth.Config) *Service {
	return &Service{
		ctx:     ctx,
		db:      db,
		spa:     spa,
		clients: make(map[string]*Client),
	}
}

func (s *Service) OmitDisplayName(omit bool) {
	s.omitDisplayName = omit
}

func (s *Service) AuthCodeURL(state string) string {
	return s.spa.AuthCodeURL(state, oauth2.AccessTypeOffline)
}

func (s *Service) Client(id primitive.ObjectID) (*Client, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	user, err := s.db.Users().FindByID(s.ctx, id)
	if err != nil {
		return nil, err
	}

	return s.clientFromUser(user)
}

func (s *Service) ClientFromSpotifyID(spotifyID string) (*Client, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	user, err := s.db.Users().FindBySpotifyID(s.ctx, spotifyID)
	if err != nil {
		return nil, err
	}

	return s.clientFromUser(user)
}

func (s *Service) Exchange(code string) (*Client, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	token, err := s.spa.Exchange(s.ctx, code)
	if err != nil {
		return nil, err
	}

	ts := s.spa.TokenSource(s.ctx, token)
	sp := s.spa.ClientWithSource(s.ctx, ts)

	me, err := sp.Me()
	if err != nil {
		return nil, err
	}

	if client, ok := s.clients[me.ID]; ok {
		s.db.Users().SaveToken(s.ctx, client.ID, token)

		ts = tokenSource(s.ctx, client.ID, s, ts)
		sp = s.spa.ClientWithSource(s.ctx, ts)
		client.replaceClient(sp)

		return client, nil
	}

	user, err := s.db.Users().FindBySpotifyID(s.ctx, me.ID)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			user = s.db.Users().Default()
			user.Token = token

			user.Spotify.ID = me.ID
			if !s.omitDisplayName {
				user.Spotify.DisplayName = me.DisplayName
			}

			if _, err := s.db.Users().Save(s.ctx, user); err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	return s.clientFromUser(user)
}

func (s *Service) clientFromUser(user *schema.User) (*Client, error) {
	if user.Token == nil {
		return nil, fmt.Errorf("Cannot create client from user with no toekn: %s", user.ID)
	}

	if client, ok := s.clients[user.Spotify.ID]; ok {
		return client, nil
	}

	ts := s.spa.TokenSource(s.ctx, user.Token)
	ts = tokenSource(s.ctx, user.ID, s, ts)
	sp := s.spa.ClientWithSource(s.ctx, ts)

	client := &Client{
		ID:      user.ID,
		User:    user.Spotify,
		service: s,
		sp:      sp,
	}
	s.clients[user.Spotify.ID] = client

	return client, nil
}
