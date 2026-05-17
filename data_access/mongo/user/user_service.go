package user

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserMongoService struct {
	collection *mongo.Collection
}

func NewUserMongoService(database *mongo.Database) *UserMongoService {
	service := &UserMongoService{
		collection: database.Collection("users"),
	}

	if err := EnsureUserIndexes(service.collection); err != nil {
		panic("failed to create user indexes: " + err.Error())
	}
	return service
}

func (mongoService *UserMongoService) UpsertUser(ctx context.Context, entity *UserEntity) error {
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

func (mongoService *UserMongoService) GetUserByEmail(ctx context.Context, email string) (*UserEntity, error) {
	var user UserEntity

	filter := bson.M{
		"email": email,
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

func (mongoService *UserMongoService) GetUserById(ctx context.Context, id string) (*UserEntity, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var user UserEntity
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
