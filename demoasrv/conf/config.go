package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"sync"

	"github.com/lee31802/comment_lib/gmicrokit"
	"github.com/lee31802/comment_lib/logkit"
	goconfig "github.com/micro/go-micro/config"
	"github.com/micro/go-micro/config/source/file"
)

type ConsumerConfig struct {
	Brokers string   `yaml:"brokers" mapstructure:"brokers"`
	Topics  []string `yaml:"topics" mapstructure:"topics"`
	Group   string   `yaml:"group" mapstructure:"group"`
	Version string   `yaml:"version" mapstructure:"version"`
}

type JobProducerConfig struct {
	Brokers string `yaml:"brokers" mapstructure:"brokers"`
	Topic   string `yaml:"topic" mapstructure:"topic"`
	Version string `yaml:"version" mapstructure:"version"`
}

type EventProducerConfig struct {
	Brokers string `yaml:"brokers" mapstructure:"brokers"`
	Topic   string `yaml:"topic" mapstructure:"topic"`
	Version string `yaml:"version" mapstructure:"version"`
}

type Kafka struct {
	RatingConsumer ConsumerConfig      `yaml:"rating_consumer"`
	JobProducer    JobProducerConfig   `yaml:"rating_producer"`
	EventProducer  EventProducerConfig `yaml:"event_producer"`
}

type RefreshUnreadStoreRatingConfig struct {
	ChannelBufferSize uint32 `yaml:"channel_buffer_size"`
	ConsumerNum       uint32 `yaml:"consumer_num"`
	Timeout           uint64 `yaml:"timeout"`
}

type RatingConfig struct {
	BaseConfig microkit.BaseConfig
	once       sync.Once
}

func (h *RatingConfig) Scan(configFile string) (err error) {
	h.once.Do(func() {

		goconfig.Load(file.NewSource(
			file.WithPath(configFile),
		))

		conf := goconfig.DefaultConfig
		err = h.BaseConfig.Scan(conf)
		if err != nil {
			logkit.Error("scan config error ", logkit.String("configFile", configFile), logkit.Err(err))
			return
		}

		//Read Values From Config if need
		var serverName string
		{
			serverName = conf.Get("server.name").String("foody-ratingsvr")

		}
		h.Load(configFile)

		// watch change in config file
		go h.watch()
		logkit.Info("Scan config finished ", logkit.String("serverName", serverName))
	})
	return
}

func (h *RatingConfig) watch() {
	w, err := goconfig.Watch("log")
	if err != nil {
		// do something
		logkit.Warn("watch config failed ", logkit.Err(err))
	}
	for {
		// wait for next value
		_, err := w.Next()

		if err != nil {
			logkit.Warn("next config failed ", logkit.Err(err))
			continue
		}
	}
}

func (h *RatingConfig) Load(configFile string) error {
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		logkit.Error("Read config error", logkit.String("filename", configFile), logkit.Err(err))
		return err
	}

	err = yaml.Unmarshal(data, h)
	if err != nil {
		logkit.Error("Unmarshl config error", logkit.String("filename", configFile), logkit.Err(err))
		return err
	}
	return nil
}
