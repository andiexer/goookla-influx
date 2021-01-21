package sinks

import (
	"fmt"
	"goookla-influx/internal/utils"
	"strconv"
	"strings"
)

type SinkSender interface {
	Send(speedtestResult *SpeedtestResult) error
}

var A string = "asdfasdf"

type SpeedtestResult struct {
	Ping float64
	Jitter float64
	Download float64
	Upload float64
}

func NewSpeedtestResult(output []byte) *SpeedtestResult {
	/*
	 *   how a speedtest result will look like
	 * 	"server name","server id","latency","jitter","packet loss","download","upload","download bytes","upload bytes","share url"
	 *  "Leucom Stafag AG - Frauenfeld","17788","5.503","0.184","0","21791359","29207386","100801072","242553115","https://www.speedtest.net/result/c/408fe7df-0d43-4ad4-bb35-8c5b2dfdcbee"
	 */
	values := strings.Split(string(output),",")
	res := &SpeedtestResult{}
	res.Ping, _ = strconv.ParseFloat(utils.RemoveQuotesFromString(values[2]), 32)
	res.Jitter, _ = strconv.ParseFloat(utils.RemoveQuotesFromString(values[3]), 32)
	download, _ := strconv.ParseFloat(utils.RemoveQuotesFromString(values[5]), 32)
	res.Download =  download / 125000
	upload, _ := strconv.ParseFloat(utils.RemoveQuotesFromString(values[6]), 32)
	res.Upload =  upload / 125000

	return res
}

func (s *SpeedtestResult) String() string {
	return fmt.Sprintf("ping=%f jitter=%f download=%f upload=%f",
		s.Ping,
		s.Jitter,
		s.Download,
		s.Upload)
}

