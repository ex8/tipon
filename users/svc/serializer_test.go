package svc

import (
	"testing"

	"github.com/ex8/tipon/users/pb"
)

func TestProtoToStruct(t *testing.T) {
	user := &pb.User{
		Username: "u",
		Password: "p",
		Address:  "a",
	}
	result := protoToStruct(user)
	if result.Username != "u" {
		t.Fatalf("serialize failed username; got %v, expected %v", result.Username, "u")
	}
}

func TestStructToProto(t *testing.T) {
	user := &user{
		Username: "u1",
		Password: "p1",
		Address:  "a1",
	}
	result := structToProto(user)
	if result.Username != "u1" {
		t.Fatalf("serialize failed username; got %v, expected %v", result.Username, "u")
	}
}
