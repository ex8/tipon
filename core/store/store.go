package store

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Store struct {
	Client *mongo.Client
}

type Opts struct {
	Host string
	Port string
}

func New(ctx context.Context, opts Opts) (*Store, error) {
	// connection string
	uri := fmt.Sprintf("mongodb://%s:%s", opts.Host, opts.Port)

	// db client connect
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	// store ping
	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, err
	}
	return &Store{Client: client}, nil
}
