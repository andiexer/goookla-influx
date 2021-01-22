package cmd

import (
	"context"
	"fmt"
	"github.com/rs/zerolog/log"
	"goookla-influx/internal/sinks"
	"os/exec"
	"runtime"
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
	args := []string{"--accept-license", "--progress=no", "--format=csv"}
	if a.serverId != "" {
		log.Info().Msgf("using serverId=%s", a.serverId)
		args = append(args, fmt.Sprintf("--server-id=%s", a.serverId))
	}
	log.Debug().Msgf("using arguments=%s", args)

	for {
		out, err := a.executeSpeedtest(args)

		if err != nil {
			log.Error().Err(err).Msg("error while exec speedtest")
		} else {
			a.sink.Send(sinks.NewSpeedtestResult(out))
		}
		log.Info().Msgf("goto sleep for seconds=%d", a.interval)
		runtime.GC()
		time.Sleep(time.Duration(a.interval) * time.Second)
	}
}

func (a App) executeSpeedtest(args []string) (string, error){
	log.Info().Msg("exec new speedtest measurement")
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, "speedtest", args...)
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}

	if ctx.Err() != nil {
		return "", ctx.Err()
	}

	return string(out), nil
}