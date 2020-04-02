package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
	"time"

	"sewik/pkg/dom/stats"
	"sewik/pkg/es"
	"sewik/pkg/sewik"
	"sewik/pkg/sys"
)

var cpuFile = flag.String("profile.cpu", "", "write cpu profile to `file`")
var memFile = flag.String("profile.mem", "", "write memory profile to `file`")
var workerNum = flag.Int("w", 5, "worker pool size")
var procNum = flag.Int("n", runtime.GOMAXPROCS(0), "set GOMAXPROCS = n")
var procDiv = flag.Int("d", 3, "set GOMAXPROCS /= d")
var command = flag.String("c", "xml", "xml|")

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

	filenames := sys.Filenames(flag.Args())
	switch *command {
	case "xml":
		printXMLStats(filenames, *workerNum)
	case "json":
		printJSON(filenames, *workerNum)
	}

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

	log.Print(sys.Stats(start, *workerNum, newProcCount, defaultProcCount))
}

func printJSON(filenames <-chan string, workerNum int) {
	fmt.Println(`{`)
	for event := range sewik.ElementsOf("ZDARZENIE", filenames, workerNum) {
		e := es.NewDoc(event)
		fmt.Print(e)
		fmt.Println(`,`)
	}
	fmt.Printf(`"__stat":"%s"}\n`, sys.MemStats())
}

func printXMLStats(filenames <-chan string, workerNum int) {
	elements := stats.NewElementsWithLock()
	for e := range sewik.RootElements(workerNum, filenames) {
		elements.Add(e)
	}
	stats.PrintXML(elements)
}
