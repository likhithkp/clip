package url

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func EnsureUrlIndexes(collection *mongo.Collection) error {
	indexes := []mongo.IndexModel{
		{
			Keys:    bson.D{{Key: "code", Value: 1}},
			Options: options.Index().SetName("code_idx"),
		},
	}

	_, err := collection.Indexes().CreateMany(context.Background(), indexes)
	return err
}
