package sys

import (
	"fmt"
	"runtime"
)

func PrintMemStats() string {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	return fmt.Sprintf(
		"\tAlloc = %s\tTotalAlloc = %s\tSys = %s\tNumGC = %v",
		bToMb(m.Alloc),
		bToMb(m.TotalAlloc),
		bToMb(m.Sys),
		m.NumGC,
	)
}

func bToMb(b uint64) string {
	if b > 1024*1024*1024 {
		return fmt.Sprintf("%v GiB", b/1024/1024/1024)
	}

	if b > 1024*1024 {
		return fmt.Sprintf("%v MiB", b/1024/1024)
	}

	if b > 1024 {
		return fmt.Sprintf("%v KiB", b/1024)
	}

	return fmt.Sprintf("%v B", b/1024)
}
