package main

import (
	"github.com/ex8/tipon/users/pb"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	r := gin.Default()

	conn, err := grpc.Dial("localhost:5000", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := pb.NewUserServiceClient(conn)

	service := userApiService{client: client}

	r.POST("/users", service.createUser)
	r.PATCH("/users", service.updateUser)

	r.Run(":5001")
}
