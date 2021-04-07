package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/ex8/tipon/core/store"
	"github.com/ex8/tipon/users/pb"
	"github.com/ex8/tipon/users/svc"
	"google.golang.org/grpc"
)

func main() {
	// config
	const (
		// servicePort = os.Getenv("TIPON_USERS_PORT")
		// mongoHost   = os.Getenv("TIPON_MONGO_HOST")
		// mongoPort   = os.Getenv("TIPON_MONGO_PORT")
		servicePort = "5000"
		mongoClient = "mongodb"
		mongoHost   = "127.0.0.1"
		mongoPort   = "27017"
	)

	// logger
	logger := log.New(os.Stdout, "USERS_SERVICE::", log.Default().Flags())

	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", servicePort))
	if err != nil {
		logger.Fatalf("failed to listen: %v", err)
	}

	// store
	s := store.New()

	// store ctx
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// store connect
	opts := store.Opts{Client: mongoClient, Host: mongoHost, Port: mongoPort}
	if err = s.Connect(ctx, opts); err != nil {
		logger.Fatalf("failed to connect to store: %v", err)
	}

	// store disconnect
	defer func() {
		if err = s.Disconnect(ctx); err != nil {
			logger.Fatalf("failed to disconnect store: %v", err)
		}
	}()

	logger.Printf("starting service on port %v\n", servicePort)

	// user grpc server
	server := grpc.NewServer()
	pb.RegisterUserServiceServer(server, &svc.UserService{Users: s.Client.Database("tipon-users").Collection("users")})
	if err := server.Serve(lis); err != nil {
		logger.Fatalf("failed to serve: %v", err)
	}
}
