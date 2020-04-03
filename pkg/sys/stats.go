package sys

import (
	"fmt"
	"log"
	"math/rand"
	"runtime"
	"time"

	"sewik/pkg/dom"
	"sewik/pkg/es"
)

func Stats(start time.Time, workerNum int, newProcCount int, defaultProcCount int) string {
	s := "Stats:\n"
	s += fmt.Sprintf("  Time: %s\n", time.Since(start))
	s += fmt.Sprintf("  Workers: %d\n", workerNum)
	s += fmt.Sprintf("  Processors: %d of %d\n", newProcCount, defaultProcCount)
	s += fmt.Sprintf("  Mem: %s\n", MemStats())

	return s
}

func MemStats() string {
	var m runtime.MemStats

	runtime.ReadMemStats(&m)

	return fmt.Sprintf(
		"Alloc = %s\tTotalAlloc = %s\tSys = %s\tNumGC = %v",
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

	return fmt.Sprintf("%v B", b)
}

var statsOn bool

func ChUtDo(s string, ch chan *es.Document) {
	if !statsOn {
		return
	}
	ticker := time.NewTicker(time.Second)
	go func(r int) {
		for range ticker.C {
			c := cap(ch)
			l := len(ch)
			u := 0
			if c > 0 {
				u = 100 * l / c
			}
			log.Printf("[%s %04d] %3d = %7d /%7d\n", s, r, u, l, c)
		}
	}(rand.Intn(9999))
}

func ChUtEl(s string, ch chan *dom.Element) {
	if !statsOn {
		return
	}
	ticker := time.NewTicker(time.Second)
	go func(r int) {
		for range ticker.C {
			c := cap(ch)
			l := len(ch)
			u := 0
			if c > 0 {
				u = 100 * l / c
			}
			log.Printf("[%s %04d] %3d = %7d /%7d\n", s, r, u, l, c)
		}
	}(rand.Intn(9999))
}
