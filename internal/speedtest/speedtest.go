package speedtest

import (
	"bytes"
	"context"
	"encoding/json"
	"os/exec"
	"time"

	"github.com/rs/zerolog/log"
)

type TestResult struct {
	Ping     Ping        `json:"ping"`
	Download Measurement `json:"download"`
	Upload   Measurement `json:"upload"`
}

type Ping struct {
	Jitter  float32 `json:"jitter"`
	Latency float32 `json:"latency"`
}

type Measurement struct {
	Bandwith int `json:"bandwidth"`
}

func (m Measurement) ToMbit() int {
	return m.Bandwith / 125000
}

var buf bytes.Buffer

func Exec() (result *TestResult, err error) {
	log.Info().Msg("exec new speedtest measurement")
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	defer buf.Reset()
	cmd := exec.Command("speedtest", "--accept-license", "--format=json")
	cmd.Stdout = &buf
	err = cmd.Run()

	if err != nil {
		return
	}
	if ctx.Err() != nil {
		return nil, ctx.Err()
	}
	result = &TestResult{}
	err = json.Unmarshal(buf.Bytes(), result)
	return
}
