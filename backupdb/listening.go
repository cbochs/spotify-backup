package backupdb

import (
	"context"

	"github.com/cbochs/spotify-backup-api/schema"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ListeningDB struct {
	db *mongo.Collection
}

func (db *ListeningDB) Default() schema.PlayHistory {
	return schema.PlayHistory{
		ID: primitive.NewObjectID(),
	}
}

func (db *ListeningDB) Save(ctx context.Context, rp []schema.PlayHistory) ([]primitive.ObjectID, error) {
	updates := make([]mongo.WriteModel, len(rp))
	for i, ph := range rp {
		updates[i] = mongo.
			NewUpdateOneModel().
			SetUpdate(bson.M{"$setOnInsert": ph}).
			SetFilter(bson.M{"user.id": ph.User.ID, "played_at": ph.PlayedAt}).
			SetUpsert(true)
	}

	options := options.BulkWrite().SetOrdered(true)

	res, err := db.db.BulkWrite(ctx, updates, options)
	if err != nil {
		return nil, err
	}

	ids := make([]primitive.ObjectID, len(rp))
	for i := range ids {
		if id, ok := res.UpsertedIDs[int64(i)]; ok {
			ids[i] = id.(primitive.ObjectID)
		}
	}

	return ids, nil
}
