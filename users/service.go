package main

import (
	"context"

	"github.com/ex8/tipon/users/pb"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type userService struct {
	pb.UnimplementedUserServiceServer
	users *mongo.Collection
}

type user struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Username string             `bson:"username"`
	Password string             `bson:"password"`
	Addredd  string             `bson:"address"`
}

func (s *userService) FindOneUser(ctx context.Context, req *pb.FindOneUserReq) (*pb.FindOneUserRes, error) {
	return nil, nil
}

func (s *userService) CreateUser(ctx context.Context, req *pb.CreateUserReq) (*pb.CreateUserRes, error) {
	return nil, nil
}

func (s *userService) UpdateUser(ctx context.Context, req *pb.UpdateUserReq) (*pb.UpdateUserRes, error) {
	return nil, nil
}
