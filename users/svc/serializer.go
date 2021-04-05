package svc

import (
	"github.com/ex8/tipon/users/pb"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func protoToStruct(u *pb.User) user {
	return user{
		Username: u.GetUsername(),
		Password: u.GetPassword(),
		Address:  u.GetAddress(),
	}
}

func structToProto(u *user) *pb.User {
	return &pb.User{
		Username: u.Username,
		Password: u.Password,
		Address:  u.Password,
	}
}

// TODO: move this to core
func objectIdToHex(id interface{}) string {
	return id.(primitive.ObjectID).Hex()
}

// TODO: move this to core
func hexToObjectId(id string) (primitive.ObjectID, error) {
	s, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return s, nil
}
