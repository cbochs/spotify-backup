package backupdb

import (
	"context"

	"github.com/cbochs/spotify-backup-api/config"
	"go.mongodb.org/mongo-driver/mongo"
)

type Client struct {
	ctx    context.Context
	config *config.DatabaseConfig
	client *mongo.Client
}

func New(ctx context.Context, config *config.DatabaseConfig) (*Client, error) {
	client, err := mongo.NewClient(config.Options())
	if err != nil {
		return nil, err
	}

	db := &Client{
		ctx:    ctx,
		config: config,
		client: client,
	}

	return db, nil
}

func (c *Client) Connect() error {
	if err := c.client.Connect(c.ctx); err != nil {
		return err
	}
	if err := c.client.Ping(c.ctx, nil); err != nil {
		return err
	}
	return nil
}

func (c *Client) Users() *UsersDB {
	coll := c.client.
		Database(c.config.Name).
		Collection("users")
	return &UsersDB{db: coll}
}

func (c *Client) Playlists() *mongo.Collection {
	return c.client.
		Database(c.config.Name).
		Collection("playlists")
}

func (c *Client) ListeningHistory() *ListeningDB {
	coll := c.client.
		Database(c.config.Name).
		Collection("listening_history")
	return &ListeningDB{db: coll}
}
