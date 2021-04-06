package svc

import (
	"github.com/ex8/tipon/tips/pb"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func protoToStruct(t *pb.Tip) tip {
	return tip{
		Amount: t.GetAmount(),
		Note:   t.GetNote(),
		UserID: t.GetUserId(),
	}
}

func structToProto(t *tip) *pb.Tip {
	return &pb.Tip{
		Id:     t.ID.Hex(),
		Amount: t.Amount,
		Note:   t.Note,
		UserId: t.UserID,
	}
}

func objectIdToHex(s interface{}) string {
	return s.(primitive.ObjectID).Hex()
}

func hexToObjectId(s string) (primitive.ObjectID, error) {
	x, err := primitive.ObjectIDFromHex(s)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return x, nil
}
