package main

import (
	"flag"
	"fmt"
	"math/rand"
	"runtime"
	"time"

	"sewik/pkg/sewik"
	"sewik/pkg/sys"
)

var (
	indexName  string
	numWorkers int
	flushBytes int
	numItems   int
	filenames  []string
)

func init() {
	flag.StringVar(&indexName, "index", "idx", "Index name")
	flag.IntVar(&numWorkers, "workers", runtime.NumCPU()/2, "Number of indexer workers")
	flag.IntVar(&flushBytes, "flush", 5e+6, "Flush threshold in bytes")
	flag.IntVar(&numItems, "count", 10000, "Number of documents to generate")

	flag.Parse()

	filenames = flag.Args()

	rand.Seed(time.Now().UnixNano())
}

func main() {
	for d := range sewik.EsDocs(sewik.ElementsOf("ZDARZENIE", sys.Filenames(filenames, 100), numWorkers, (numItems+1)*numWorkers)) {
		fmt.Println(d.String())
	}
}
