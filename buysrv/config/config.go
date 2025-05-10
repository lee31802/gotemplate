package config

import (
	"github.com/lee31802/comment_lib/gzero/server"
)

var Cfg *Config

// Config 整合所有配置部分
type Config struct {
	Log    LogConfig `mapstructure:"log"`
	Server BuyConfig `mapstructure:"server"`
}

// EtcdConfig 对应配置中的 buy.Etcd 部分
type EtcdConfig struct {
	Hosts              []string
	Key                string
	ID                 int64  `json:",optional"`
	User               string `json:",optional"`
	Pass               string `json:",optional"`
	CertFile           string `json:",optional"`
	CertKeyFile        string `json:",optional=CertFile"`
	CACertFile         string `json:",optional=CertFile"`
	InsecureSkipVerify bool   `json:",optional"`
}

type ServerMiddlewaresConf struct {
	Trace      bool `json:",default=true"`
	Recover    bool `json:",default=true"`
	Stat       bool `json:",default=true"`
	Prometheus bool `json:",default=true"`
	Breaker    bool `json:",default=true"`
}

// BuyConfig 对应配置中的 buy 部分
type BuyConfig struct {
	ServiceConf struct {
		Name string
	}
	ListenOn      string
	Etcd          EtcdConfig `json:",optional,inherit"`
	Auth          bool       `json:",optional"`
	StrictControl bool       `json:",optional"`
	// setting 0 means no timeout
	Timeout      int64 `json:",default=2000"`
	CpuThreshold int64 `json:",default=900,range=[0:1000]"`
	// grpc health check switch
	Health      bool `json:",default=true"`
	Middlewares ServerMiddlewaresConf
}

// LogConfig 对应配置中的 log 部分
type LogConfig struct {
	Level         string
	Path          string
	MaxSize       int
	MaxBackups    int
	BufferSize    int
	ChannelSize   int
	MaxAge        int
	EnableCaller  bool
	EnableConsole bool
	ErrorAsync    bool
}

func InitCfg() error {
	return server.Config.Unmarshal(&Cfg)
}
