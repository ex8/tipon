package main

import (
	"log"
	"net"

	"github.com/ex8/tipon/tips/pb"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterTipServiceServer(s, &tipService{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to service: %v", err)
	}
}
