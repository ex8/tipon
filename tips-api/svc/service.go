package svc

import (
	"io"
	"net/http"

	"github.com/ex8/tipon/tips/pb"
	"github.com/gin-gonic/gin"
)

type TipApiService struct {
	TipClient pb.TipServiceClient
}

type tip struct {
	ID     string  `json:"id,omitempty"`
	Amount float64 `json:"amount" binding:"required"`
	Note   string  `json:"note"`
	UserID string  `json:"user_id" binding:"required"`
}

func (s *TipApiService) GetTips(ctx *gin.Context) {
	// call stream to find tips from tips grpc service
	stream, err := s.TipClient.FindTips(ctx, &pb.FindTipsReq{})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	for {
		// receive tip from stream
		tip, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// TODO: build streaming api ?
		ctx.JSON(http.StatusOK, gin.H{"tip": tip})
		return
	}
}

func (s *TipApiService) GetTipById(ctx *gin.Context) {
	// get id from params
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "empty id param"})
		return
	}

	// call find one tips grpc service
	res, err := s.TipClient.FindOneTip(ctx, &pb.FindOneTipReq{Id: id})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"tip": res.GetTip()})
}

func (s *TipApiService) CreateTip(ctx *gin.Context) {
	// bind tip
	var t tip
	if err := ctx.ShouldBindJSON(&t); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// serialize data (struct -> pb)
	tip := &pb.Tip{
		Amount: t.Amount,
		Note:   t.Note,
		UserId: t.UserID,
	}

	// call create tip grpc service
	res, err := s.TipClient.CreateTip(ctx, &pb.CreateTipReq{Tip: tip})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"tip": res.GetTip()})
}
