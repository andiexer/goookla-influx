package cmd

import (
	"github.com/rs/zerolog/log"
	"goookla-influx/internal/sinks"
	"goookla-influx/internal/speedtest"
	"time"
)

type App struct {
	sink sinks.SinkSender
	interval uint
	serverId string
}

func NewApp(sink sinks.SinkSender, interval uint, serverId string) *App {
	return &App{sink: sink, interval: interval, serverId: serverId}
}

func (a *App) Run() {
	log.Info().Msg("starting gookla speedtest app")
	ticker := time.NewTicker(time.Duration(a.interval) * time.Second).C
	for {
		select {
			case <- ticker:
				out, err := speedtest.Exec()
				if err != nil {
					log.Error().Err(err).Msg("error while exec speedtest")
				} else {
					a.sink.Send(out)
				}
		}
	}
}