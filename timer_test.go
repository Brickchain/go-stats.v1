package stats_test

import stats "github.com/Brickchain/go-stats.v1"

func ExampleStartTimer() {
	t := stats.StartTimer("api.get.total")
	defer t.Stop()
}
