package sinks

import (
	"github.com/rs/zerolog/log"
	"github.com/influxdata/influxdb1-client/v2"
	"time"
)

type InfluxDbV1Sink struct {
	influxUrl string
	user string
	password string
}

func NewInfluxDbV1Sink(influxUrl string, user string, password string) *InfluxDbV1Sink {
	return &InfluxDbV1Sink{influxUrl: influxUrl, user: user, password: password}
}

func (i *InfluxDbV1Sink) Send(speedtestResult *SpeedtestResult) error {
	log.Debug().Msg("sending data to influxdb v1")
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr: i.influxUrl,
		Username: i.user,
		Password: i.password,
	})

	if err != nil {
		return err
	}
	defer c.Close()

	bp, _ := client.NewBatchPoints(client.BatchPointsConfig{
		Database: "speedtest",
		Precision: "s",
	})

	speedtest_fields := map[string]interface{}{
		"ping": speedtestResult.Ping,
		"jitter" : speedtestResult.Jitter,
		"dowload": speedtestResult.Download,
		"upload" : speedtestResult.Upload,
	}
	pt, err := client.NewPoint("speedtest", nil, speedtest_fields, time.Now())
	if err != nil {
		return err
	}
	bp.AddPoint(pt)
	return c.Write(bp)
}