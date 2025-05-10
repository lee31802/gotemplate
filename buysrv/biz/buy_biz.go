package biz

import "github.com/lee31802/gotemplate/buysrv/service"

type BuyBiz struct {
	buyService *service.BuyService
}

func NewBuyBiz(buyService *service.BuyService) *BuyBiz {
	return &BuyBiz{
		buyService: buyService,
	}
}
