package main

import (
	microkit "github.com/lee31802/comment_lib/gmicrokit"
	"github.com/lee31802/comment_lib/gweb"
	"github.com/lee31802/comment_lib/logkit"
	"github.com/lee31802/gotemplate/gateway/demoapi/ping"
)

var (
	cmd = &gweb.Command{
		Name: "test",
		PreRun: func(router gweb.Router) error {
			router.GET("/", func() string { return "OK" })
			return nil
		},
		Modules: []gweb.Module{&ping.PingModule{}},
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
	microkit.Init()
	cmd.Execute()
}
