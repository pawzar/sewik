package xml

import (
	"fmt"
	"log"
	"path/filepath"
	"sync"

	"github.com/subchen/go-xmldom"

	"sewik/sys"
	"sewik/xml/parse"
	"sewik/xml/print"
	"sewik/xml/structure"
)

func ScanXMLsInPaths(p []string, workerPoolSize int) {
	filenames := make(chan string)
	nodes := make(chan *xmldom.Node)

	var wg sync.WaitGroup
	go func() {
		wg.Wait()
		log.Println("[OUT] EOD")
		close(nodes)
	}()

	for w := 1; w <= workerPoolSize; w++ {
		wg.Add(1)
		go worker(&wg, w, filenames, nodes)
	}

	go populateFilenames(filenames, p)

	elems := structure.NewElementsWithLock()
	for node := range nodes {
		elems.Add(node)
	}

	print.Elements(elems, 0, 0)

	fmt.Println()
}

func populateFilenames(filenames chan<- string, patterns []string) {
	defer close(filenames)

	for k, pattern := range patterns {
		log.Printf("[IN] - %d -", k)

		files, _ := filepath.Glob(pattern)
		for _, file := range files {
			filenames <- file
			log.Printf("[IN] %q", file)
		}
	}

	log.Println("[IN] DONE")
}

func worker(wg *sync.WaitGroup, n int, jobs <-chan string, results chan<- *xmldom.Node) {
	defer wg.Done()

	log.Printf("[WORKER] %d START", n)

	for filename := range jobs {
		log.Printf("[WORKER] %d [BEGINS] %q\n%s", n, filename, sys.PrintMemStats())

		scanFile(results, filename)

		log.Printf("[WORKER] %d [FINISHED] %q\n%s", n, filename, sys.PrintMemStats())
	}

	log.Printf("[WORKER] %d STOP", n)
}

func scanFile(nodes chan<- *xmldom.Node, filename string) {
	fmt.Printf("<!-- %s -->\n", filename)
	doc, err := parse.File(filename)
	if err != nil {
		log.Println(filename + ": " + err.Error())

		return
	}

	nodes <- doc.Root
}
