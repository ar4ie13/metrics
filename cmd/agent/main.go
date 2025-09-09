package main

import (
	"github.com/ar4ie13/metrics/internal/agent/config"
	"github.com/ar4ie13/metrics/internal/agent/handler"
	"github.com/ar4ie13/metrics/internal/agent/service"
	"log"
	"time"
)

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}

}

func run() error {

	cfg := config.NewAgentConfig()
	cfg.InitAgentConfig()
	ms := service.NewMetricsStorage()
	hndlr := handler.NewHandler(ms, cfg)
	pollInterval := cfg.GetPollInterval()
	reportInterval := cfg.GetReportInterval()

	collectTicker := time.NewTicker(time.Duration(pollInterval) * time.Second)
	sendTicker := time.NewTicker(time.Duration(reportInterval) * time.Second)
	log.Printf("Server endpoint: %s\nPoll interval: %d\nReport interval: %d\n\n", cfg.GetEndpointServerAddr(), pollInterval, reportInterval)
	for {
		select {
		case <-collectTicker.C:
			log.Println("Collecting metrics")
			ms.UpdateMetrics()
		case <-sendTicker.C:
			log.Println("Sending metrics")
			hndlr.PostMetrics()
		}
	}

}
