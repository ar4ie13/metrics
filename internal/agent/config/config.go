package config

import (
	"flag"
)

type AgentConfig struct {
	endpointServerAddr string
	pollInterval       int
	reportInterval     int
}

func NewAgentConfig() *AgentConfig {
	return &AgentConfig{}
}

func (c *AgentConfig) InitAgentConfig() {
	flag.StringVar(&c.endpointServerAddr, "a", "localhost:8080", "server address endpoint")
	flag.IntVar(&c.pollInterval, "p", 2, "poll interval (sec)")
	flag.IntVar(&c.reportInterval, "r", 10, "report interval (sec)")
	flag.Parse()
}

func (c *AgentConfig) GetEndpointServerAddr() string {
	return c.endpointServerAddr
}

func (c *AgentConfig) GetPollInterval() int {
	return c.pollInterval
}

func (c *AgentConfig) GetReportInterval() int {
	return c.reportInterval
}
