package sinks

import (
	"github.com/rs/zerolog/log"
)

type ConsoleSink struct {}

func NewConsoleSink() *ConsoleSink {
	return &ConsoleSink{}
}

func (c *ConsoleSink) Send(speedtestResult *SpeedtestResult) error {
	log.Info().Msg("=== speedtest results ===")
	log.Info().Msg(speedtestResult.String())
	log.Info().Msg("=== speedtest results ===")
	return nil
}
