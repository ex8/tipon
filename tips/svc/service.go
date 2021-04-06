package svc

import (
	"context"

	"github.com/ex8/tipon/tips/pb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TipService struct {
	pb.UnimplementedTipServiceServer
	Tips *mongo.Collection
}

type tip struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	Amount float64            `bson:"amount"`
	Note   string             `bson:"note"`
	UserID string             `bson:"user_id"`
}

func (s *TipService) FindTips(req *pb.FindTipsReq, stream pb.TipService_FindTipsServer) error {
	ctx := context.Background()

	// find
	cursor, err := s.Tips.Find(ctx, &bson.M{})
	if err != nil {
		return err
	}
	defer cursor.Close(ctx)

	// paginate
	for cursor.Next(ctx) {
		// decode data
		data := tip{}
		if err := cursor.Decode(&data); err != nil {
			return err
		}

		// serialize data (struct -> pb)
		t := structToProto(&data)

		// send data
		if err := stream.Send(&pb.FindTipsRes{Tip: t}); err != nil {
			return err
		}
	}
	if err := cursor.Err(); err != nil {
		return err
	}
	return nil
}

func (s *TipService) FindOneTip(ctx context.Context, req *pb.FindOneTipReq) (*pb.FindOneTipRes, error) {
	// convert hex to object id
	id, err := hexToObjectId(req.GetId())
	if err != nil {
		return nil, err
	}

	// find one
	result := s.Tips.FindOne(ctx, bson.M{"_id": id})

	// decode data
	data := tip{}
	if err := result.Decode(&data); err != nil {
		return nil, err
	}

	// serialize data (struct -> pb)
	t := structToProto(&data)
	return &pb.FindOneTipRes{Tip: t}, nil
}

func (s *TipService) CreateTip(ctx context.Context, req *pb.CreateTipReq) (*pb.CreateTipRes, error) {
	t := req.GetTip()

	// serialize data (pb -> struct)
	data := protoToStruct(t)

	// insert data
	result, err := s.Tips.InsertOne(ctx, data)
	if err != nil {
		return nil, err
	}

	// assign & convert object id to hex
	t.Id = objectIdToHex(result.InsertedID)
	return &pb.CreateTipRes{Tip: t}, nil
}
