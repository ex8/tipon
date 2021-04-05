package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/ex8/tipon/users/pb"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"google.golang.org/grpc"
)

func main() {
	var (
		servicePort = os.Getenv("TIPON_USERS_PORT")
		mongoHost   = os.Getenv("TIPON_MONGO_HOST")
		mongoPort   = os.Getenv("TIPON_MONGO_PORT")
	)

	logger := log.New(os.Stdout, "USERS_SERVICE::", log.Default().Flags())

	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", servicePort))
	if err != nil {
		logger.Fatalf("users failed to listen: %v", err)
	}

	// DB connect
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	uri := fmt.Sprintf("mongodb://%s:%s", mongoHost, mongoPort)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		logger.Fatalf("users failed db connection: %v", err)
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	// DB ping
	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		logger.Fatalf("users failed db ping: %v", err)
	}

	logger.Printf("starting users service on port %v\n", servicePort)

	// User gRPC server
	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &userService{users: client.Database("tipon-users").Collection("users")})
	if err := s.Serve(lis); err != nil {
		logger.Fatalf("users failed to serve: %v", err)
	}
}
