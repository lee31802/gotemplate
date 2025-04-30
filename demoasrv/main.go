package main

import (
	"flag"
	"github.com/lee31802/comment_lib/env"
	microkit "github.com/lee31802/comment_lib/gmicrokit"
	"github.com/lee31802/comment_lib/logkit"
	"os"
)

var (
	//path = "ratingsvr/conf/rating_test_tw_conf.yml"
	configFile = flag.String("config", "conf/conf.yml", "config file")
)

func main() {
	flag.Parse()

	// 环境变量 is_local=true 会让服务只注册在本地
	isLocal := os.Getenv("is_local")
	if isLocal == "true" {
		microkit.Init(microkit.RegistryReadonly(true))
		logkit.Info("microkit init from local environment")
	} else {
		microkit.Init()
	}

	svr := &DemoSvr{}
	svr.Init(*configFile)

	if env.GetEnv() != "live" {
		logkit.SetLevel("debug")
	}
	svr.Run()
}
