package main

import (
	"context"

	"github.com/ex8/tipon/tips/pb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type tipService struct {
	pb.UnimplementedTipServiceServer
	tips *mongo.Collection
}

type tip struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	Amount float64            `bson:"amount"`
	Note   string             `bson:"note"`
	UserID string             `bson:"user_id"`
}

func (s *tipService) FindTips(req *pb.FindTipsReq, stream pb.TipService_FindTipsServer) error {
	ctx := context.Background()
	cursor, err := s.tips.Find(ctx, &bson.M{})
	if err != nil {
		return err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		data := tip{}
		if err := cursor.Decode(&data); err != nil {
			return err
		}
		t := &pb.Tip{
			Id:     data.ID.Hex(),
			Amount: data.Amount,
			Note:   data.Note,
			UserId: data.UserID,
		}
		if err := stream.Send(&pb.FindTipsRes{Tip: t}); err != nil {
			return err
		}
	}
	if err := cursor.Err(); err != nil {
		return err
	}
	return nil
}

func (s *tipService) FindOneTip(ctx context.Context, req *pb.FindOneTipReq) (*pb.FindOneTipRes, error) {
	id, err := primitive.ObjectIDFromHex(req.GetId())
	if err != nil {
		return nil, err
	}
	result := s.tips.FindOne(ctx, bson.M{"_id": id})
	data := tip{}
	if err := result.Decode(&data); err != nil {
		return nil, err
	}
	t := &pb.Tip{
		Id:     data.ID.Hex(),
		Amount: data.Amount,
		Note:   data.Note,
		UserId: data.UserID,
	}
	return &pb.FindOneTipRes{Tip: t}, nil
}

func (s *tipService) CreateTip(ctx context.Context, req *pb.CreateTipReq) (*pb.CreateTipRes, error) {
	t := req.GetTip()
	data := tip{
		Amount: t.GetAmount(),
		Note:   t.GetNote(),
		UserID: t.GetUserId(),
	}
	result, err := s.tips.InsertOne(ctx, data)
	if err != nil {
		return nil, err
	}
	t.Id = result.InsertedID.(primitive.ObjectID).Hex()
	return &pb.CreateTipRes{Tip: t}, nil
}
