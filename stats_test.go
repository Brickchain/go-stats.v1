package stats_test

import stats "github.com/Brickchain/go-stats.v1"

func ExampleSetup_inmem() {
	// The inmem sink prints the collected metrics to stdout once per minute.
	stats.Setup("inmem", "example")
}

func ExampleSetup_datadog() {
	// Requires that the DOGSTATSD environment variable is set and pointing to the statsd port of your dd-agent.
	// export DOGSTATSD="localhost:8125"
	stats.Setup("datadog", "example")
}

func ExampleSetup_prometheus() {
	// The prometheus sink adds the /metrics handler to your HTTP stack.
	// You need to setup a HTTP listener on some port that you point Prometheus to.
	stats.Setup("prometheus", "example")
}
