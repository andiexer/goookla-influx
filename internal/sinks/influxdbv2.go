package sinks

import (
	"context"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/rs/zerolog/log"
	"goookla-influx/internal/speedtest"
)

type InfluxDbV2Sink struct {
	influxUrl string
	authToken string
	orgUnit string
	bucket string
}

func NewInfluxDbV2Sink(influxUrl string, authToken string, orgUnit string, bucket string) *InfluxDbV2Sink {
	return &InfluxDbV2Sink{influxUrl: influxUrl, authToken: authToken, orgUnit: orgUnit, bucket: bucket}
}

func (i *InfluxDbV2Sink) Send(result *speedtest.TestResult) (err error) {
	log.Debug().Msg("sending data to influxdb v2")
	client := influxdb2.NewClient(i.influxUrl, i.authToken)
	defer client.Close()
	writeApi := client.WriteAPIBlocking(i.orgUnit,i.bucket)

	p := influxdb2.NewPointWithMeasurement("speedtest").
		AddField(pingField, result.Ping.Latency).
		AddField(jitterField, result.Ping.Jitter).
		AddField(downloadField, result.Download.ToMbit()).
		AddField(uploadField, result.Upload.ToMbit())

	err = writeApi.WritePoint(context.Background(), p)
	return
}