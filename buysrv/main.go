package main

import (
	"flag"
	"github.com/lee31802/comment_lib/gzero/server"
	"github.com/lee31802/comment_lib/logkit"
	"github.com/lee31802/gotemplate/buysrv/config"
	"github.com/lee31802/gotemplate/buysrv/wire"
	"github.com/lee31802/gotemplate/pb/buy"
	"google.golang.org/grpc"
)

var (
	cmd = &server.Command{
		PreRun: func() error {
			err := logkit.Init()
			if err != nil {
				panic(err)
			}
			err = config.InitCfg()
			if err != nil {
				panic(err)
			}
			logkit.Info("init config success")
			return nil
		},
		RegisterServer: func(grpcServer *grpc.Server) {
			buy.RegisterBuyServer(grpcServer, wire.InitializeHandler())
		},
	}
)

func main() {
	var appPath = flag.String("appPath", "", "Application path")
	cmd.AppPath = *appPath
	if err := cmd.Execute(); err != nil {
		logkit.With(logkit.Err(err)).Error("run buy server error")
	}
}
