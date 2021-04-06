package main

import (
	"fmt"
	"log"

	"github.com/ex8/tipon/users-api/svc"
	"github.com/ex8/tipon/users/pb"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	// config
	const (
		servicePort     = "5001"
		userServiceName = "localhost"
		userServicePort = "5000"
	)

	// default router
	r := gin.Default()

	// user client connection
	host := fmt.Sprintf("%v:%v", userServiceName, userServicePort)
	conn, err := grpc.Dial(host, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("unable to establish client connection to %v - %v", host, err)
	}
	defer conn.Close()

	// user grpc client
	client := pb.NewUserServiceClient(conn)

	// new user api service
	s := svc.UserApiService{UserClient: client}

	// routes
	r.POST("/users", s.CreateUser)
	r.PATCH("/users/:id", s.UpdateUser)

	// serve
	r.Run(fmt.Sprintf(":%v", servicePort))
}
