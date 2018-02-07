package stats_test

import (
	stats "github.com/Brickchain/go-stats.v1"
)

func ExampleIncrement() {
	stats.Increment("api.get.calls", 1)
}

func ExampleGauge() {
	response := []byte{"Hello, World"}
	stats.Gauge("api.get.bytes", len(response))
}
