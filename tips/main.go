package main

import (
	"context"
	"log"
	"net"
	"time"

	"github.com/ex8/tipon/tips/pb"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalf("tips failed to listen: %v", err)
	}

	// DB connect
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatalf("tips failed db connection: %v", err)
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	// DB ping
	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatalf("tips failed db ping: %v", err)
	}

	// Tips gRPC server
	s := grpc.NewServer()
	pb.RegisterTipServiceServer(s, &tipService{tips: client.Database("tipon-tips").Collection("tips")})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("tips failed to serve: %v", err)
	}
}
