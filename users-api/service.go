package main

import (
	"net/http"

	"github.com/ex8/tipon/users/pb"
	"github.com/gin-gonic/gin"
)

type user struct {
	ID       string `json:"id,omitempty"`
	Username string `json:"username"`
	Password string `json:"password"`
	Address  string `json:"address"`
}

type userApiService struct {
	client pb.UserServiceClient
}

func (s *userApiService) createUser(ctx *gin.Context) {
	var u user
	if err := ctx.ShouldBindJSON(&u); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user := &pb.User{
		Username: u.Username,
		Password: u.Password,
		Address:  u.Address,
	}
	res, err := s.client.CreateUser(ctx, &pb.CreateUserReq{User: user})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"user": res.GetUser()})
}

func (s *userApiService) updateUser(ctx *gin.Context) {
	var u user
	if err := ctx.ShouldBindJSON(&u); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user := &pb.User{
		Id:       u.ID,
		Username: u.Username,
		Password: u.Password,
		Address:  u.Address,
	}
	res, err := s.client.UpdateUser(ctx, &pb.UpdateUserReq{User: user})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"user": res.GetUser()})
}
