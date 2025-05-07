package main

import (
	"github.com/lee31802/comment_lib/gzero/server"
	"github.com/lee31802/gotemplate/buysrv/config"
	"github.com/lee31802/gotemplate/buysrv/handler"
	"github.com/lee31802/gotemplate/pb/buy"
	"google.golang.org/grpc"
)

func NewServer() *BuySvr {
	return &BuySvr{}
}

var (
	cmd = &server.Command{
		RegisterFunc: func(grpcServer *grpc.Server) {
			buy.RegisterBuyServer(grpcServer, handler.NewDemoHandler())
		},
	}
)

type BuySvr struct {
	config  config.RatingConfig
	handler handler.DemoHandler
}

func (h *BuySvr) Init(conf string) error {
	//err := h.config.Scan(conf)
	//if err != nil {
	//	return err
	//}
	//err = cmd.Init(h.GetBaseConfig())
	//if err != nil {
	//	return err
	//}
	//
	//// foody_rating
	//h.service = buy.NewService()
	////ratingHandler := wire.InitializeHandler(&h.config.Mysql, &h.config.Redis,
	////	&h.config.Kafka.JobProducer, &h.config.Kafka.EventProducer, &h.config.Spex, &h.config.RefreshUnreadStoreRatingConfig)
	//return buy.RegisterDemoHandler(&handler.DemoHandler{})
	return nil
}

func (h *BuySvr) Run() {
	cmd.Execute()
	//h.service.Run()
}

//func InitLogkit(conf *config.RecentVisitConfig) {
//	if conf.Log.Level == "" {
//		conf.Log.Level = logkit.LOG_LEVEL
//	}
//	if conf.Log.Path == "" {
//		conf.Log.Path = logkit.LOG_PATH
//	}
//	if conf.Log.MaxSize == 0 {
//		conf.Log.MaxSize = logkit.LOG_MAX_SIZE
//	}
//	if conf.Log.MaxBackups == 0 {
//		conf.Log.MaxBackups = logkit.LOG_MAX_BACKUPS
//	}
//	if conf.Log.MaxAge == 0 {
//		conf.Log.MaxAge = logkit.LOG_MAX_AGE
//	}
//	err := logkit.Init(
//		logkit.Level(conf.Log.Level),
//		logkit.Path(conf.Log.Path),
//		logkit.MaxAge(conf.Log.MaxAge),
//		logkit.MaxBackups(conf.Log.MaxBackups),
//		logkit.MaxSize(conf.Log.MaxSize),
//		logkit.EnableConsole(conf.Log.EnableConsole),
//		logkit.EnableCaller(conf.Log.EnableCaller),
//		logkit.ErrorAsync(conf.Log.ErrorAsync),
//	)
//	if err != nil {
//		panic(err)
//	}
//}
