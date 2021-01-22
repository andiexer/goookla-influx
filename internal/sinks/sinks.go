package sinks

import "goookla-influx/internal/speedtest"

type SinkSender interface {
	Send(result *speedtest.TestResult) (err error)
}

const pingField = "ping"
const jitterField = "jitter"
const downloadField = "download"
const uploadField = "upload"

