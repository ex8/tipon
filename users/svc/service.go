package svc

import (
	"context"

	"github.com/ex8/tipon/users/pb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
	Users *mongo.Collection
}

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Username string             `bson:"username" json:"username"`
	Password string             `bson:"password" json:"password"`
	Address  string             `bson:"address" json:"address"`
}

func (s *UserService) CreateUser(ctx context.Context, req *pb.CreateUserReq) (*pb.CreateUserRes, error) {
	u := req.GetUser()

	// serialize data (pb -> struct)
	data := protoToStruct(u)

	// insert data
	result, err := s.Users.InsertOne(ctx, data)
	if err != nil {
		return nil, err
	}

	// assign & convert object id to hex
	u.Id = objectIdToHex(result.InsertedID)
	return &pb.CreateUserRes{User: u}, nil
}

func (s *UserService) UpdateUser(ctx context.Context, req *pb.UpdateUserReq) (*pb.UpdateUserRes, error) {
	u := req.GetUser()

	// convert hex to object id
	id, err := hexToObjectId(u.GetId())
	if err != nil {
		return nil, err
	}

	// serialize data (pb -> struct)
	data := protoToStruct(u)

	// SetReturnDocument ensures updated doc is returned
	opts := options.FindOneAndUpdate().SetReturnDocument(1)

	// find by id and update only changed fields using $set
	result := s.Users.FindOneAndUpdate(ctx, &bson.M{"_id": id}, bson.M{"$set": data}, opts)
	if result.Err() != nil {
		return nil, result.Err()
	}

	// decode updated result
	decoded := &User{}
	if err := result.Decode(&decoded); err != nil {
		return nil, err
	}

	// serialize updated data (struct -> pb)
	updated := structToProto(decoded)
	return &pb.UpdateUserRes{User: updated}, nil
}
