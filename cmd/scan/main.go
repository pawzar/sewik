package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
	"time"

	"sewik/pkg/dom"
	"sewik/pkg/dom/stats"
	"sewik/pkg/es"
	"sewik/pkg/sewik"
	"sewik/pkg/sys"
)

var cpuFile = flag.String("profile.cpu", "", "write cpu profile to `file`")
var memFile = flag.String("profile.mem", "", "write memory profile to `file`")
var workNum = flag.Int("w", 5, "worker pool size")
var pipeSize = flag.Int("p", 10000, "pipe size per one worker")
var procNum = flag.Int("n", runtime.GOMAXPROCS(0), "set GOMAXPROCS = n")
var procDiv = flag.Int("d", 3, "set GOMAXPROCS /= d")
var cmd = flag.String("c", "xml", "xml|json")

func main() {
	start := time.Now()
	flags()
	newProcCount := *procNum / *procDiv
	defaultProcCount := runtime.GOMAXPROCS(newProcCount)
	cpuStats(*cpuFile)
	commands(*cmd, *workNum, *pipeSize)
	memStats(*memFile)
	log.Print(sys.Stats(start, *workNum, newProcCount, defaultProcCount))
}

func printJSON(filenames <-chan string, workerNum int, pipeSize int) {
	for event := range sewik.ElementsOf("ZDARZENIE", filenames, workerNum, workerNum*(pipeSize+1)) {
		e := es.NewDoc(event)
		fmt.Println(e)
	}
}

func printJSON2(filenames <-chan string, workerNum int, pipeSize int) {
	ch := sewik.ElementsOf("ZDARZENIE", filenames, workerNum, workerNum*(pipeSize+1))
	//cnt := dom.NewRollingCounter()
	for nn := range ch {
		//counter := dom.NewCounter().WithNode(nn)
		//cnt.Add(counter)
		fmt.Println(dom.NewObject().From(nn).String())
	}
}

func printXMLStats(filenames <-chan string, workerNum int, pipeSize int) {
	elements := stats.NewElementsWithLock()
	for e := range sewik.ElementsOf("ZDARZENIE", filenames, workerNum, workerNum*(pipeSize+1)) {
		elements.Add(e)
	}
	stats.PrintXML(elements)
}

func commands(s string, workerCount int, pipeSize int) {
	filenames := sys.Filenames(flag.Args(), 500)
	switch s {
	case "xml":
		printXMLStats(filenames, workerCount, pipeSize)
	case "json":
		printJSON(filenames, workerCount, pipeSize)
	case "c":
		printJSON2(filenames, workerCount, pipeSize)
	}
}

func flags() {
	flag.Parse()
	if flag.NArg() == 0 {
		fmt.Fprintln(os.Stderr, "Usage: sewik (files... | \"glob\") [options]\nOptions:")
		flag.PrintDefaults()
		os.Exit(1)
	}
}

func cpuStats(s string) {
	if s != "" {
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
}

func memStats(s string) {
	if s != "" {
		f, err := os.Create(*memFile)
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		defer f.Close()
		runtime.GC()
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
	}
}
