package svc

import "github.com/ex8/tipon/users/pb"

func structToProto(u *user) *pb.User {
	return &pb.User{
		Id:       u.ID,
		Username: u.Username,
		Password: u.Password,
		Address:  u.Address,
	}
}
