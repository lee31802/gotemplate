package client

import (
	"github.com/lee31802/comment_lib/gzero/client"
	"github.com/lee31802/gotemplate/pb/buy"
)

var (
	demoAClient buy.Client
)

func DemoAClient() buy.Client {
	if demoAClient == nil {
		demoAClient = buy.NewClient(client.InitClient())
	}
	return demoAClient
}
