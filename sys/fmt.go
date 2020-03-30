package sys

import (
	"fmt"
	"runtime"
	"time"
)

func FormattedStats(start time.Time, workerNum int, newProcCount int, defaultProcCount int) string {
	s := fmt.Sprintf("Time: %s\n", time.Since(start)) +
		fmt.Sprintf("Workers: %d\n", workerNum) +
		fmt.Sprintf("Processors: %d of %d\n", newProcCount, defaultProcCount) +
		fmt.Sprintf("Mem: \n%s\n", PrintMemStats())
	runtime.GC()
	s += fmt.Sprintf("Post GC Mem: \n%s\n", PrintMemStats())

	return s
}
