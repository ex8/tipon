package main

import (
	"context"

	"github.com/ex8/tipon/tips/pb"
)

type tipService struct {
	pb.UnimplementedTipServiceServer
}

func (s *tipService) FindTips(req *pb.FindTipsRequest, stream pb.TipService_FindTipsServer) error {
	return nil
}

func (s *tipService) FindOneTip(ctx context.Context, req *pb.FindOneTipRequest) (*pb.Tip, error) {
	return &pb.Tip{Amount: 5.0, Note: "f", UserId: "fd"}, nil
}

func (s *tipService) CreateTip(ctx context.Context, req *pb.CreateTipRequest) (*pb.Tip, error) {
	return &pb.Tip{Amount: 5.0, Note: "f", UserId: "fd"}, nil
}

func (s *tipService) UpdateTip(ctx context.Context, req *pb.UpdateTipRequest) (*pb.Tip, error) {
	return &pb.Tip{Amount: 5.0, Note: "f", UserId: "fd"}, nil
}
