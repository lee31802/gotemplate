package main

import (
	"flag"
	"github.com/lee31802/comment_lib/gweb"
	"github.com/lee31802/comment_lib/logkit"
	"github.com/lee31802/gotemplate/gateway/demoapi/buy"
	"github.com/lee31802/gotemplate/gateway/demoapi/client"
	"github.com/lee31802/gotemplate/gateway/demoapi/config"
	"github.com/lee31802/gotemplate/gateway/demoapi/ping"
)

var (
	cmd = &gweb.Command{
		Name: "demo",
		PreRun: func(router gweb.Router) error {
			err := logkit.Init()
			if err != nil {
				panic(err)
			}
			err = config.InitCfg()
			if err != nil {
				panic(err)
			}
			logkit.Info("init config success")
			router.GET("/", func() string { return "OK" })
			return nil
		},
		Modules: []gweb.Module{&ping.PingModule{}, buy.NewDemoModule()},
	}
)

func main() {
	var appPath = flag.String("appPath", "", "Application path")
	cmd.AppPath = *appPath
	client.InitClient()
	if err := cmd.Execute(); err != nil {
		logkit.With(logkit.Err(err)).Error("run application error")
	}
}
