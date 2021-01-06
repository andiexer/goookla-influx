package cmd

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"goookla-influx/internal/sinks"
	"os/exec"
	"time"
)

type App struct {
	sink *sinks.SinkSender
	interval uint
	serverId string
}

func NewApp(sink *sinks.SinkSender, interval uint, serverId string) *App {
	return &App{sink: sink, interval: interval, serverId: serverId}
}

func (a *App) Run() {
	log.Info().Msg("starting gookla speedtest app")
	args := []string{"--progress=no", "--format=csv"}
	if a.serverId != "" {
		log.Info().Msgf("using serverId=%s", a.serverId)
		args = append(args, fmt.Sprintf("--server-id=%s", a.serverId))
	}
	log.Debug().Msgf("using arguments=%s", args)

	for true {
		log.Info().Msg("exec new speedtest measurement")

		out, err := exec.Command("speedtest",args...).Output()
		if err != nil {
			log.Err(err)
		}

		res := sinks.NewSpeedtestResult(out)
		(*a.sink).Send(res)

		log.Info().Msgf("goto sleep for seconds=%d", a.interval)
		time.Sleep(time.Duration(a.interval) * time.Second)
	}

}

