package main

import "github.com/lee31802/comment_lib/ginserver"

var (
	cmd = &ginserver.Command{
		Name: "test",
		PreRun: func(router ginserver.Router) error {
			router.GET("/", func() string { return "OK" })
			return nil
		},
		Modules: []ginserver.Module{},
	}
)
