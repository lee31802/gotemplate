package main

import (
	"flag"
	"github.com/lee31802/comment_lib/gweb"
	"github.com/lee31802/comment_lib/logkit"
	"github.com/lee31802/gotemplate/gateway/demoapi/buy"
	"github.com/lee31802/gotemplate/gateway/demoapi/ping"
)

var (
	cmd = &gweb.Command{
		Name: "test",
		PreRun: func(router gweb.Router) error {
			router.GET("/", func() string { return "OK" })
			return nil
		},
		Modules: []gweb.Module{&ping.PingModule{}, buy.NewDemoModule()},
	}
)

func InitLogkit() {
	logkit.Init(
		logkit.Level("debug"),
		logkit.Path("log"),
		logkit.MaxSize(1024),
		logkit.MaxBackups(1024),
		logkit.MaxAge(1),
		logkit.EnableCaller(true),
		logkit.EnableConsole(true),
		logkit.ErrorAsync(false),
	)
}

func main() {
	InitLogkit()
	var appPath = flag.String("appPath", "", "Application path")
	//var configFile = flag.String("config", "conf/config.yml", "config file")
	cmd.AppPath = *appPath
	cmd.Execute()
}
