package main

import (
	"fmt"
	"log"
	"time"

	"github.com/ar4ie13/metrics/internal/agent/config"
	"github.com/ar4ie13/metrics/internal/agent/requests"
	"github.com/ar4ie13/metrics/internal/agent/service"
)

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}

}

func run() error {

	cfg := config.NewAgentConfig()
	srv := service.NewService()
	hndlr := requests.NewRequests(srv, *cfg)

	sendMetrics(*cfg, *srv, *hndlr)

	return nil
}

func sendMetrics(cfg config.AgentConfig, srv service.Service, hndlr requests.Requests) {
	collectTicker := time.NewTicker(time.Duration(cfg.PollInterval) * time.Second)
	sendTicker := time.NewTicker(time.Duration(cfg.ReportInterval) * time.Second)
	log.Printf("Server endpoint: %s\nPoll interval: %d\nReport interval: %d\n\n", cfg.EndpointServerAddr,
		cfg.PollInterval, cfg.ReportInterval)
	for {
		select {
		case <-collectTicker.C:
			log.Println("Collecting metrics")

			srv.UpdateMetrics()

			if a, ok := srv.Metrics["PollCount"]; ok {
				fmt.Println(*a.Delta)
			}

		case <-sendTicker.C:
			log.Println("Sending metrics")

			hndlr.PostMetrics()
			srv.ResetPollCount()
		}
	}
}
