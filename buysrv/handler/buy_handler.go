package handler

import (
	"context"
	"fmt"
	"github.com/lee31802/gotemplate/buysrv/biz"
	"github.com/lee31802/gotemplate/pb/buy"
)

type BuyHandler struct {
	buy.UnimplementedBuyServer
	buyBiz *biz.BuyBiz
}

func NewBuyHandler(buyBiz *biz.BuyBiz) *BuyHandler {
	return &BuyHandler{
		buyBiz: buyBiz,
	}
}
func (s *BuyHandler) CreateOrder(ctx context.Context, req *buy.CreateOrderRequest) (*buy.CreateOrderResponse, error) {
	orderId := fmt.Sprintf("id:%v", *req.UserId)
	return &buy.CreateOrderResponse{
		OrderId: &orderId,
	}, nil
}
