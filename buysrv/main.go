package main

import (
	"flag"
	"github.com/lee31802/comment_lib/env"
	"github.com/lee31802/comment_lib/logkit"
)

var (
	//path = "ratingsvr/conf/rating_test_tw_conf.yml"
	configFile = flag.String("config", "conf/conf.yml", "config file")
)

func main() {
	flag.Parse()

	svr := &BuySvr{}
	svr.Init(*configFile)

	if env.Environment() != "live" {
		logkit.SetLevel("debug")
	}
	svr.Run()
}
