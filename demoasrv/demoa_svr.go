package main

import (
	microkit "github.com/lee31802/comment_lib/gmicrokit"
	"github.com/lee31802/comment_lib/logkit"
	config "github.com/lee31802/gotemplate/demoasrv/conf"
	"github.com/lee31802/gotemplate/demoasrv/handler"
	"github.com/lee31802/gotemplate/pb/demoa"
	"github.com/micro/go-micro"
)

var (
	cmd = &microkit.Command{
		Name: "test",
		PreRun: func() error {

			return nil
		},
	}
)

type DemoSvr struct {
	config  config.RatingConfig
	service micro.Service
}

func (h *DemoSvr) GetBaseConfig() *microkit.BaseConfig {
	return &h.config.BaseConfig
}

func (h *DemoSvr) Init(conf string) error {
	err := h.config.Scan(conf)
	if err != nil {
		return err
	}
	err = cmd.Init(h.GetBaseConfig())
	if err != nil {
		return err
	}

	// foody_rating
	h.service = demoa.NewService()
	//ratingHandler := wire.InitializeHandler(&h.config.Mysql, &h.config.Redis,
	//	&h.config.Kafka.JobProducer, &h.config.Kafka.EventProducer, &h.config.Spex, &h.config.RefreshUnreadStoreRatingConfig)
	return demoa.RegisterDemoHandler(&handler.DemoHandler{})
}

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

func (h *DemoSvr) Run() {
	h.service.Run()
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
