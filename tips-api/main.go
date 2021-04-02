package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.GET("/tips", getTips)
	r.GET("/tips/:id", getTipById)
	r.POST("/tips", createTip)

	r.Run(":8001")
}
