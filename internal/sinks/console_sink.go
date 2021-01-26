package sinks

import (
	"goookla-influx/internal/speedtest"

	"github.com/rs/zerolog/log"
)

type ConsoleSink struct{}

func NewConsoleSink() *ConsoleSink {
	return &ConsoleSink{}
}

func (c *ConsoleSink) Send(result *speedtest.TestResult) error {
	log.Info().
		Float32("ping", result.Ping.Latency).
		Float32("jitter", result.Ping.Jitter).
		Int("download", result.Download.ToMbit()).
		Int("upload", result.Upload.ToMbit()).
		Msgf("ping=%v, jitter=%v,download=%v,upload=%v", result.Ping.Latency, result.Ping.Jitter, result.Download.Bandwith, result.Upload.Bandwith)
	return nil
}
