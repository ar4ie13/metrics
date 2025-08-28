package config

import (
	"flag"
)

type Config struct {
	LocalServerAddr string
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) InitConfig() {
	flag.StringVar(&c.LocalServerAddr, "a", "localhost:8080", "local server address")
	flag.Parse()
}

func (c *Config) GetLocalServerAddr() string {
	return c.LocalServerAddr
}
