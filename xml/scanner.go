package xml

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"

	"sewik/sys"
	"sewik/xml/dom"
	"sewik/xml/printer"
	"sewik/xml/spec"
)

func ScanInPaths(p []string, workerPoolSize int) {
	filenames := make(chan string)
	nodes := make(chan *dom.Element)

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

	go populateJobs(filenames, p)

	elements := spec.NewElementsWithLock()
	for node := range nodes {
		elements.Add(node)
	}

	printer.PrintElements(elements)
}

func populateJobs(filenames chan<- string, patterns []string) {
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

func worker(wg *sync.WaitGroup, n int, jobs <-chan string, results chan<- *dom.Element) {
	defer wg.Done()

	log.Printf("[WORKER] %d START", n)

	for filename := range jobs {
		log.Printf("[WORKER] %d [BEGINS] %q\n%s", n, filename, sys.MemStats())

		scan(filename, results)

		log.Printf("[WORKER] %d [FINISHED] %q\n%s", n, filename, sys.MemStats())
	}

	log.Printf("[WORKER] %d STOP", n)
}

func scan(filename string, results chan<- *dom.Element) {
	fmt.Printf("<!-- %s -->\n", filename)

	doc, err := parse(filename)
	if err != nil {
		log.Println(filename + ": " + err.Error())

		return
	}

	results <- doc.Root
}

func parse(filename string) (*dom.Document, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return dom.Parse(file)
}
