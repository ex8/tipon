package main

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	r := gin.Default()

	conn, err := grpc.Dial("localhost:8000", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	r.GET("/tips", getTips)
	r.GET("/tips/:id", getTipById)
	r.POST("/tips", createTip)

	r.Run(":8001")
}
