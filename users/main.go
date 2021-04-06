package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/ex8/tipon/users/pb"
	"github.com/ex8/tipon/users/svc"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"google.golang.org/grpc"
)

func main() {
	// config
	const (
		// servicePort = os.Getenv("TIPON_USERS_PORT")
		// mongoHost   = os.Getenv("TIPON_MONGO_HOST")
		// mongoPort   = os.Getenv("TIPON_MONGO_PORT")
		servicePort = "5000"
		mongoHost   = "127.0.0.1"
		mongoPort   = "27017"
	)

	// logger
	logger := log.New(os.Stdout, "USERS_SERVICE::", log.Default().Flags())

	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", servicePort))
	if err != nil {
		logger.Fatalf("users failed to listen: %v", err)
	}

	// store
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

	// store
	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		logger.Fatalf("users failed db ping: %v", err)
	}

	logger.Printf("starting users service on port %v\n", servicePort)

	// user grpc server
	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &svc.UserService{Users: client.Database("tipon-users").Collection("users")})
	if err := s.Serve(lis); err != nil {
		logger.Fatalf("users failed to serve: %v", err)
	}
}
