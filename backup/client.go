package backup

import (
	"context"

	"github.com/cbochs/spotify-backup-api/schema"
	"github.com/cbochs/spotify-backup-api/spotify"
	"github.com/cbochs/spotify-backup-api/spotify/options"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Client struct {
	ID      primitive.ObjectID
	User    schema.SpotifyUser
	service *Service
	sp      *spotify.Client
}

func NewClient(id primitive.ObjectID, user schema.SpotifyUser, s *Service, sp *spotify.Client) *Client {
	return &Client{
		ID:      id,
		service: s,
		sp:      sp,
	}
}

func (c *Client) replaceClient(sp *spotify.Client) {
	c.sp = sp
}

func (c *Client) BackupRecentlyPlayed(ctx context.Context) error {
	cur, err := c.sp.RecentlyPlayedOpt(options.Query().Limit(50))
	if err != nil {
		return err
	}

	srp := cur.Items
	for {
		if cur.Next() == "" {
			break
		}
		if err := c.sp.Next(cur); err != nil {
			return err
		}
		srp = append(srp, cur.Items...)
	}

	rp := make([]schema.PlayHistory, len(srp))
	for i, ph := range srp {
		track := c.service.db.ListeningHistory().Default()
		track.User = c.User
		track.PlayedAt = ph.PlayedAt.Time

		artists := make([]schema.BasicArtist, len(ph.Track.Artists))
		for i, artist := range ph.Track.Artists {
			artists[i] = schema.BasicArtist{
				Name: artist.Name,
				ID:   artist.ID,
			}
		}

		track.Track = schema.BasicTrack{
			Name:    ph.Track.Name,
			ID:      ph.Track.ID,
			Artists: artists,
		}

		rp[i] = track
	}

	if _, err := c.service.db.ListeningHistory().Save(ctx, rp); err != nil {
		return err
	}

	if err := c.service.db.Users().Touch(ctx, c.ID); err != nil {
		return err
	}

	return nil
}
