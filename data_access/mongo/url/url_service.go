package url

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UrlMongoService struct {
	collection *mongo.Collection
}

func NewUrlMongoService(database *mongo.Database) *UrlMongoService {
	service := &UrlMongoService{
		collection: database.Collection("urls"),
	}

	if err := EnsureUrlIndexes(service.collection); err != nil {
		panic("failed to create user indexes: " + err.Error())
	}
	return service
}

func (mongoService *UrlMongoService) UpsertUrl(ctx context.Context, entity *UrlEntity) error {
	filter := bson.M{"_id": entity.Id}
	update := bson.M{
		"$set": entity,
	}
	opts := options.Update().SetUpsert(true)
	_, err := mongoService.collection.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		return err
	}

	return nil
}

func (mongoService *UrlMongoService) GetUrlByCode(ctx context.Context, code string) (*UrlEntity, error) {
	var user UrlEntity

	filter := bson.M{
		"code": code,
	}

	err := mongoService.collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (mongoService *UrlMongoService) GetUserById(ctx context.Context, id string) (*UrlEntity, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var user UrlEntity
	filter := bson.M{
		"_id": oid,
	}

	err = mongoService.collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}
