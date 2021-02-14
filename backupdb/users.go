package backupdb

import (
	"context"
	"time"

	"github.com/cbochs/spotify-backup-api/schema"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/oauth2"
)

type UsersDB struct {
	db *mongo.Collection
}

func (db *UsersDB) Default() *schema.User {
	now := time.Now()
	return &schema.User{
		ID:          primitive.NewObjectID(),
		Created:     now,
		LastUpdated: now,
	}
}

func (db *UsersDB) Touch(ctx context.Context, id primitive.ObjectID) error {
	_, err := db.db.UpdateOne(
		ctx,
		bson.M{"_id": id},
		bson.M{"$set": bson.M{"last_updates": time.Now()}},
	)
	if err != nil {
		return err
	}
	return nil
}

func (db *UsersDB) FindByID(ctx context.Context, id primitive.ObjectID) (*schema.User, error) {
	res := db.db.FindOne(ctx, bson.M{"_id": id})
	if res.Err() != nil {
		return nil, res.Err()
	}

	var user schema.User
	if err := res.Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (db *UsersDB) FindBySpotifyID(ctx context.Context, spotifyID string) (*schema.User, error) {
	res := db.db.FindOne(ctx, bson.M{"spotify.id": spotifyID})
	if res.Err() != nil {
		return nil, res.Err()
	}

	var user schema.User
	if err := res.Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (db *UsersDB) Save(ctx context.Context, user *schema.User) (primitive.ObjectID, error) {
	if user.Spotify.ID == "" {
		return primitive.NilObjectID, errors.New("Spotify user ID must not be empty")
	}

	var id primitive.ObjectID
	if user.ID.IsZero() {
		res, err := db.db.InsertOne(ctx, user)
		if err != nil {
			return primitive.NilObjectID, err
		}

		id = res.InsertedID.(primitive.ObjectID)
	} else {
		res, err := db.db.ReplaceOne(ctx, bson.M{"_id": user.ID}, user)
		if err != nil {
			return primitive.NilObjectID, err
		}

		id = res.UpsertedID.(primitive.ObjectID)
	}

	return id, nil
}

func (db *UsersDB) SaveToken(ctx context.Context, id primitive.ObjectID, token *oauth2.Token) error {
	_, err := db.db.UpdateOne(
		ctx,
		bson.M{"_id": id},
		bson.M{"$set": bson.M{"token": token}},
	)
	if err != nil {
		return err
	}
	return nil
}
