package config

import (
	"os"

	"github.com/kelseyhightower/envconfig"
	"github.com/rs/zerolog/log"
)

type AppConfig struct {
	Config *Config
}

type Config struct {
	Sink            string `envconfig:"SINK" default:"console"`
	ServerId		string `envconfig:"SERVER_ID" default:""`
	InfluxHost      string `envconfig:"INFLUX_HOST" default:"http://localhost:8086"`
	InfluxDatabase	string `envconfig:"INFLUX_V1_DATABASE" default:"speedtest"`
	InfluxUser      string `envconfig:"INFLUX_V1_USER" default:"admin"`
	InfluxPassword  string `envconfig:"INFLUX_V1_PASSWORD" default:"admin"`
	InfluxOrg       string `envconfig:"INFLUX_V2_ORG" default:"devlabs"`
	InfluxBucket    string `envconfig:"INFLUX_V2_BUCKET" default:"speedtest"`
	InfluxAuthToken string `envconfig:"INFLUX_V2_AUTHTOKEN" default:""`
	Interval        uint   `envconfig:"INTERVAL" default:"900"`
}

func LoadConfig() *AppConfig {
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal().Err(err).Msg("unable to load configuration")
		os.Exit(-1)
	}

	return &AppConfig{
		Config: &cfg,
	}
}
