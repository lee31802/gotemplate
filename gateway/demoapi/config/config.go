package config

import (
	"github.com/lee31802/comment_lib/gweb"
	"time"
)

var Cfg *Config

// GinWebConfig 对应配置中的 ginweb 部分
type GinWebConfig struct {
	Recovery bool         `mapstructure:"recovery"`
	Pprof    bool         `mapstructure:"pprof"`
	Jaeger   JaegerConfig `mapstructure:"jaeger"`
}

// JaegerConfig 对应配置中的 ginweb.jaeger 部分
type JaegerConfig struct {
	Enable       bool    `mapstructure:"enable"`
	SamplingRate float64 `mapstructure:"sampling_rate"`
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

// Config 整合所有配置部分
type Config struct {
	GinWeb GinWebConfig  `mapstructure:"ginweb"`
	Log    LogConfig     `mapstructure:"log"`
	Buy    RpcClientConf `mapstructure:"buy"`
}

type RpcClientConf struct {
	Etcd          EtcdConfig    `json:",optional,inherit"`
	Endpoints     []string      `json:",optional"`
	Target        string        `json:",optional"`
	App           string        `json:",optional"`
	Token         string        `json:",optional"`
	NonBlock      bool          `json:",optional"`
	Timeout       int64         `json:",default=2000"`
	KeepaliveTime time.Duration `json:",default=20s"`
}

func InitCfg() error {
	return gweb.Config.Unmarshal(&Cfg)
}
