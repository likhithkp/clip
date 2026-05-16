package mongo

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
)

type Client = mongo.Client
type SessionContent = mongo.SessionContext

func Transaction[T any](ctx context.Context, client *mongo.Client, fn func(ctx mongo.SessionContext) (T, error)) (T, error) {
	var zero T
	session, err := client.StartSession()
	if err != nil {
		return zero, err
	}
	defer session.EndSession(ctx)

	result, err := session.WithTransaction(ctx, func(sc mongo.SessionContext) (interface{}, error) {
		return fn(sc)
	})

	if err != nil {
		return zero, err
	}
	res, ok := result.(T)
	if !ok {
		return zero, fmt.Errorf("type assertion to %T failed", zero)
	}
	return res, nil
}
