package handler

import (
	"context"
	"fmt"
	"github.com/lee31802/gotemplate/pb/buy"
)

type DemoHandler struct {
	buy.UnimplementedBuyServer
}

func NewDemoHandler() *DemoHandler {
	return &DemoHandler{}
}
func (s *DemoHandler) CreateOrder(ctx context.Context, req *buy.CreateOrderRequest) (*buy.CreateOrderResponse, error) {
	orderId := fmt.Sprintf("id:%v", *req.UserId)
	return &buy.CreateOrderResponse{
		OrderId: &orderId,
	}, nil
}
