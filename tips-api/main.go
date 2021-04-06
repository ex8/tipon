package main

import (
	"fmt"
	"log"

	"github.com/ex8/tipon/tips-api/svc"
	"github.com/ex8/tipon/tips/pb"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	// config
	const (
		servicePort     = "8001"
		tipsServiceName = "localhost"
		tipsServicePort = "8000"
	)

	// router
	r := gin.Default()

	// tips client connection
	host := fmt.Sprintf("%v:%v", tipsServiceName, tipsServicePort)
	conn, err := grpc.Dial(host, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("unable to establish client connection to %v - %v", host, err)
	}
	defer conn.Close()

	// tips grpc client
	client := pb.NewTipServiceClient(conn)

	// new tips api service
	s := svc.TipApiService{TipClient: client}

	// routes
	r.GET("/tips", s.GetTips)
	r.GET("/tips/:id", s.GetTipById)
	r.POST("/tips", s.CreateTip)

	// serve
	r.Run(fmt.Sprintf(":%v", servicePort))
}
