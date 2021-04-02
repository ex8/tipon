package main

import "github.com/gin-gonic/gin"

func getTips(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"k": "find",
	})
}

func getTipById(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"k": ctx.Param("id"),
	})
}

func createTip(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"k": "create",
	})
}
