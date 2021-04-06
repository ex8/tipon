package svc

import "testing"

func TestStructToProto(t *testing.T) {
	user := &user{
		ID:       "i1",
		Username: "u1",
		Password: "p1",
		Address:  "a1",
	}
	result := structToProto(user)
	if result.Id != "i1" {
		t.Fatalf("serialize failed username; got %v, expected %v", result.Id, "i1")
	}
}
