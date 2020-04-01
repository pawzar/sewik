package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
	"time"

	"sewik/app/scanner"
	"sewik/sys"
)

var cpuFile = flag.String("profile.cpu", "", "write cpu profile to `file`")
var memFile = flag.String("profile.mem", "", "write memory profile to `file`")
var pool = flag.Int("w", 5, "worker pool size")
var procNum = flag.Int("n", runtime.GOMAXPROCS(0), "set GOMAXPROCS = n")
var procDiv = flag.Int("d", 3, "set GOMAXPROCS /= d")

func main() {
	start := time.Now()

	flag.Parse()

	if flag.NArg() == 0 {
		fmt.Fprintln(os.Stderr, "Usage: sewik (files... | \"glob\") [options]\nOptions:")
		flag.PrintDefaults()
		os.Exit(1)
	}

	newProcCount := *procNum / *procDiv
	defaultProcCount := runtime.GOMAXPROCS(newProcCount)
	log.Printf("procs (max: %d): %d", defaultProcCount, newProcCount)

	if *cpuFile != "" {
		f, err := os.Create(*cpuFile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		defer f.Close() // error handling omitted for example
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}

	workerNum := *pool

	scanner.ScanFilesInPaths(flag.Args(), workerNum)
	//fmt.Print(`{`)
	//for event := range sewik.EventChannel(flag.Args(), workerNum) {
	//	e := es.NewDoc(event)
	//	fmt.Print(e)
	//	fmt.Println(`,`)
	//}
	//fmt.Printf(`"__stat":"%s"}`, sys.MemStats())

	if *memFile != "" {
		f, err := os.Create(*memFile)
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		defer f.Close() // error handling omitted for example
		runtime.GC()    // get up-to-date statistics
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
	}

	log.Print(sys.Stats(start, workerNum, newProcCount, defaultProcCount))
}
