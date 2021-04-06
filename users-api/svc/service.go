package svc

import (
	"net/http"

	"github.com/ex8/tipon/users/pb"
	"github.com/gin-gonic/gin"
)

type UserApiService struct {
	UserClient pb.UserServiceClient
}

type user struct {
	ID       string `json:"id,omitempty"`
	Username string `json:"username"`
	Password string `json:"password"`
	Address  string `json:"address"`
}

func (s *UserApiService) CreateUser(ctx *gin.Context) {
	// bind user
	var u user
	if err := ctx.ShouldBindJSON(&u); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// serialize data (struct -> pb)
	user := structToProto(&u)

	// call create user grpc service
	res, err := s.UserClient.CreateUser(ctx, &pb.CreateUserReq{User: user})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"user": res.GetUser()})
}

func (s *UserApiService) UpdateUser(ctx *gin.Context) {
	// bind user
	var u user
	if err := ctx.ShouldBindJSON(&u); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// assign user id from params
	u.ID = ctx.Param("id")

	// serialize data (struct -> pb)
	user := structToProto(&u)

	// call update user grpc service
	res, err := s.UserClient.UpdateUser(ctx, &pb.UpdateUserReq{User: user})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"user": res.GetUser()})
}
