package store

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type StoreService interface {
	Connect(ctx context.Context, opts Opts) error
	Disconnect(ctx context.Context) error
}

type Store struct {
	Client *mongo.Client
}

type Opts struct {
	Client string
	Host   string
	Port   string
}

func (s *Store) Connect(ctx context.Context, opts Opts) error {
	uri := fmt.Sprintf("%s://%s:%s", opts.Client, opts.Host, opts.Port)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return err
	}
	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		return err
	}
	s.SetClient(ctx, client)
	return nil
}

func (s *Store) Disconnect(ctx context.Context) error {
	if err := s.Client.Disconnect(ctx); err != nil {
		return err
	}
	return nil
}

func (s *Store) SetClient(ctx context.Context, client *mongo.Client) {
	s.Client = client
}

func New(ctx context.Context, opts Opts) *Store {
	return &Store{}
}
