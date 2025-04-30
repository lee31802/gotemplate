package client

import (
	"github.com/lee31802/gotemplate/pb/demoa"
)

var (
	demoAClient demoa.DemoClient
)

func DemoAClient() demoa.DemoClient {
	if demoAClient == nil {
		initDemoAClient()
	}
	return demoAClient
}

func initDemoAClient() {
	cli := demoa.NewClient()

	demoAClient = cli

}
