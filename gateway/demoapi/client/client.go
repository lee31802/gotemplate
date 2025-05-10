package client

import (
	"github.com/lee31802/comment_lib/gzero/client"
	"github.com/lee31802/gotemplate/pb/buy"
)

var (
	buyClient buy.Client
)

func InitClient() {
	BuyClient()
}

func BuyClient() buy.Client {
	if buyClient == nil {
		buyClient = buy.NewClient(client.InitClient(client.WithServiceName(buy.GetServiceName())))
	}
	return buyClient
}
