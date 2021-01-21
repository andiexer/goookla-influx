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
	database string
}

func NewInfluxDbV1Sink(influxUrl string, user string, password string, database string) *InfluxDbV1Sink {
	return &InfluxDbV1Sink{influxUrl: influxUrl, user: user, password: password, database: database}
}

func (i *InfluxDbV1Sink) Send(speedtestResult *SpeedtestResult) (err error) {
	log.Debug().Msg("sending data to influxdb v1")
	var c client.Client
	c, err = client.NewHTTPClient(client.HTTPConfig{
		Addr: i.influxUrl,
		Username: i.user,
		Password: i.password,
	})

	if err != nil {
		return
	}
	defer c.Close()

	bp, _ := client.NewBatchPoints(client.BatchPointsConfig{
		Database: i.database,
		Precision: "s",
	})

	speedtest_fields := map[string]interface{}{
		pingField: speedtestResult.Ping,
		jitterField : speedtestResult.Jitter,
		downloadField: speedtestResult.Download,
		uploadField : speedtestResult.Upload,
	}
	var pt *client.Point
	pt, err = client.NewPoint("speedtest", nil, speedtest_fields, time.Now())
	if err != nil {
		return
	}
	bp.AddPoint(pt)
	return c.Write(bp)
}