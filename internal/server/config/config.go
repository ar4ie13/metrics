package config

import (
	"flag"

	handlerConfig "github.com/ar4ie13/metrics/internal/server/handler/config"
)

type Config struct {
	HandlerConfig handlerConfig.Config
}

func NewConfig() *Config {
	cfg := &Config{}
	cfg.InitConfig()
	return cfg
}

func (c *Config) InitConfig() {
	defaultLocalServerAddr := "localhost:8080"
	flag.StringVar(&c.HandlerConfig.LocalServerAddr, "a", defaultLocalServerAddr, "local server address")
	flag.Parse()
}
