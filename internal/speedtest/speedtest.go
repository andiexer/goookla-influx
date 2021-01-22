package speedtest

import (
	"context"
	"encoding/json"
	"github.com/rs/zerolog/log"
	"os/exec"
	"time"
)

type TestResult struct {
	Ping Ping `json:"ping"`
	Download Measurement `json:"download"`
	Upload Measurement `json:"upload"`
}

type Ping struct {
	Jitter float32 		`json:"jitter"`
	Latency float32 	`json:"latency"`
}

type Measurement struct {
	Bandwith int 		`json:"bandwith"`
}

func (m Measurement) ToMbit() float32 {
	return float32(m.Bandwith)
}

func Exec() (result *TestResult, err error) {
	log.Info().Msg("exec new speedtest measurement")
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	cmd := exec.Command("speedtest", "--accept-license", "--format=json")
	out, err := cmd.Output()

	if err != nil {
		return
	}
	if ctx.Err() != nil {
		return nil, ctx.Err()
	}

	result = &TestResult{}
	err = json.Unmarshal(out, result)
	return
}