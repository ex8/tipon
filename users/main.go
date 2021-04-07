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
		mongoHost   = "127.0.0.1"
		mongoPort   = "27017"
	)

	// logger
	logger := log.New(os.Stdout, "USERS_SERVICE::", log.Default().Flags())

	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", servicePort))
	if err != nil {
		logger.Fatalf("failed to listen: %v", err)
	}

	// Store ctx
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db, err := store.New(ctx, store.Opts{Host: mongoHost, Port: mongoPort})
	if err != nil {
		logger.Fatalf("failed to connect db: %v", err)
	}
	defer func() {
		if err = db.Disconnect(ctx); err != nil {
			logger.Fatalf("failed to disconnect db: %v", err)
		}
	}()

	logger.Printf("starting service on port %v\n", servicePort)

	// user grpc server
	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &svc.UserService{Users: db.Client.Database("tipon-users").Collection("users")})
	if err := s.Serve(lis); err != nil {
		logger.Fatalf("failed to serve: %v", err)
	}
}
