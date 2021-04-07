package store

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Store *mongo.Client

type StoreOpts struct {
	Host string
	Port string
}

func NewStore(ctx context.Context, opts StoreOpts) (Store, error) {
	// connection string
	uri := fmt.Sprintf("mongodb://%s:%s", opts.Host, opts.Port)

	// store (db) client connect
	store, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	// store ping
	if err = store.Ping(ctx, readpref.Primary()); err != nil {
		return nil, err
	}
	return store, nil
}
