//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/lee31802/gotemplate/buysrv/biz"
	"github.com/lee31802/gotemplate/buysrv/dao"
	"github.com/lee31802/gotemplate/buysrv/handler"
	"github.com/lee31802/gotemplate/buysrv/service"
)

func InitializeHandler() *handler.BuyHandler {
	wire.Build(
		handler.NewBuyHandler,
		biz.NewBuyBiz,
		service.NewBuyService,
		dao.NewBuyDao,
	)
	return &handler.BuyHandler{}
}
