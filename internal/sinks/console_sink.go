package sinks

import (
	"github.com/rs/zerolog/log"
	"goookla-influx/internal/speedtest"
)

type ConsoleSink struct {}

func NewConsoleSink() *ConsoleSink {
	return &ConsoleSink{}
}

func (c *ConsoleSink) Send(result *speedtest.TestResult) error {
	log.Info().Msg("=== speedtest results ===")
	log.Info().Msgf("ping=%v, jitter=%v,download=%v,upload=%v", result.Ping.Latency, result.Ping.Jitter, result.Download.Bandwith, result.Upload.Bandwith)
	log.Info().Msg("=== speedtest results ===")
	return nil
}
