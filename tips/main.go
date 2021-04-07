package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/ex8/tipon/core/store"
	"github.com/ex8/tipon/tips/pb"
	"github.com/ex8/tipon/tips/svc"
	"google.golang.org/grpc"
)

func main() {
	var (
		// servicePort = os.Getenv("TIPON_TIPS_PORT")
		// mongoHost   = os.Getenv("TIPON_MONGO_HOST")
		// mongoPort   = os.Getenv("TIPON_MONGO_PORT")
		servicePort = "8000"
		mongoClient = "mongodb"
		mongoHost   = "127.0.0.1"
		mongoPort   = "27017"
	)

	logger := log.New(os.Stdout, "TIPS_SERVICE::", log.Default().Flags())

	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", servicePort))
	if err != nil {
		logger.Fatalf("tips failed to listen: %v", err)
	}

	// store
	s := store.New()

	// store ctx
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// store connect
	opts := store.Opts{Client: mongoClient, Host: mongoHost, Port: mongoPort}
	if err = s.Connect(ctx, opts); err != nil {
		logger.Fatalf("failed to connect store: %v", err)
	}

	// store disconnect
	defer func() {
		if err = s.Disconnect(ctx); err != nil {
			logger.Fatalf("failed to disconnected store: %v", err)
		}
	}()

	logger.Printf("starting tips service on port %v\n", servicePort)

	// Tips gRPC server
	server := grpc.NewServer()
	pb.RegisterTipServiceServer(server, &svc.TipService{Tips: s.Client.Database("tipon-tips").Collection("tips")})
	if err := server.Serve(lis); err != nil {
		logger.Fatalf("tips failed to serve: %v", err)
	}
}
