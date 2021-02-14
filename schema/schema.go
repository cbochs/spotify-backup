package schema

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/oauth2"
)

type User struct {
	ID          primitive.ObjectID   `bson:"_id"`
	Spotify     SpotifyUser          `bson:"spotify"`
	Token       *oauth2.Token        `bson:"token"`
	Playlists   []primitive.ObjectID `bson:"playlists"`
	LikedSongs  *primitive.ObjectID  `bson:"liked_songs"`
	Created     time.Time            `bson:"created"`
	LastUpdated time.Time            `bson:"last_updated"`
}

type SpotifyUser struct {
	ID          string `bson:"id"`
	DisplayName string `bson:"display_name,omitempty"`
}

type PlayHistory struct {
	ID       primitive.ObjectID `bson:"_id"`
	User     SpotifyUser        `bson:"user"`
	PlayedAt time.Time          `bson:"played_at"`
	Track    BasicTrack         `bson:"track"`
}

type BasicTrack struct {
	Name    string        `bson:"name"`
	ID      string        `bson:"id"`
	Artists []BasicArtist `bson:"artists"`
}

type BasicArtist struct {
	Name string `bson:"name"`
	ID   string `bson:"id"`
}
