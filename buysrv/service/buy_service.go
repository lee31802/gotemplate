package service

import "github.com/lee31802/gotemplate/buysrv/dao"

type BuyService struct {
	buyDao *dao.BuyDao
}

func NewBuyService(buyDao *dao.BuyDao) *BuyService {
	return &BuyService{
		buyDao: buyDao,
	}
}
