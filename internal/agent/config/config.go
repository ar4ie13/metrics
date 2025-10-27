package config

import (
	"flag"
)

type AgentConfig struct {
	EndpointServerAddr string
	PollInterval       int
	ReportInterval     int
}

func NewAgentConfig() *AgentConfig {
	cfg := &AgentConfig{}
	cfg.InitAgentConfig()
	return cfg
}

func (c *AgentConfig) InitAgentConfig() {
	flag.StringVar(&c.EndpointServerAddr, "a", "localhost:8080", "server address endpoint")
	flag.IntVar(&c.PollInterval, "p", 2, "poll interval (sec)")
	flag.IntVar(&c.ReportInterval, "r", 10, "report interval (sec)")
	flag.Parse()
}
