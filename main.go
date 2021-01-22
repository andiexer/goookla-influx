package main

import (
	"github.com/rs/zerolog/log"
	"goookla-influx/cmd"
	"goookla-influx/internal/config"
	"goookla-influx/internal/sinks"
	"strings"
)

func main() {
	//defer profile.Start(profile.MemProfile, profile.MemProfileRate(1), profile.ProfilePath(".")).Stop()
	cfg := config.LoadConfig()
	var sink sinks.SinkSender
	switch strings.ToLower(cfg.Config.Sink) {
	case "console":
		log.Debug().Msg("using console sink")
		sink = sinks.NewConsoleSink()
	case "influxdbv2":
		log.Debug().Msg("using influxdb v2")
		sink = sinks.NewInfluxDbV2Sink(cfg.Config.InfluxHost, cfg.Config.InfluxAuthToken, cfg.Config.InfluxOrg, cfg.Config.InfluxBucket)
	case "influxdbv1":
		log.Debug().Msg("using influxdb v1")
		sink = sinks.NewInfluxDbV1Sink(cfg.Config.InfluxHost, cfg.Config.InfluxUser, cfg.Config.InfluxPassword, cfg.Config.InfluxDatabase)
	default:
		log.Fatal().Msgf("the provided sink=%s is not possible", cfg.Config.Sink)
	}
	app := cmd.NewApp(sink, cfg.Config.Interval, cfg.Config.ServerId)
	app.Run()
}
