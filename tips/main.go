package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/ex8/tipon/tips/pb"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"google.golang.org/grpc"
)

func main() {
	var (
		servicePort = os.Getenv("TIPON_TIPS_PORT")
		mongoHost   = os.Getenv("TIPON_MONGO_HOST")
		mongoPort   = os.Getenv("TIPON_MONGO_PORT")
	)

	logger := log.New(os.Stdout, "tips_server:", log.Default().Flags())

	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", servicePort))
	if err != nil {
		logger.Fatalf("tips failed to listen: %v", err)
	}

	// DB connect
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	uri := fmt.Sprintf("mongodb://%s:%s", mongoHost, mongoPort)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		logger.Fatalf("tips failed db connection: %v", err)
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	// DB ping
	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		logger.Fatalf("tips failed db ping: %v", err)
	}

	// Tips gRPC server
	s := grpc.NewServer()
	pb.RegisterTipServiceServer(s, &tipService{tips: client.Database("tipon-tips").Collection("tips")})
	if err := s.Serve(lis); err != nil {
		logger.Fatalf("tips failed to serve: %v", err)
	}
}
