package sewik

import (
	"log"
	"os"
	"path/filepath"
	"sync"

	"sewik/dom"
	"sewik/sys"
	"sewik/xml"
)

func newWorkerPool() *workerPool {
	return &workerPool{
		waiter:    &sync.WaitGroup{},
		filenames: make(stringChan),
		events:    make(elementChan),
	}
}

type stringChan chan string
type elementChan chan *dom.Element
type workerPool struct {
	waiter    *sync.WaitGroup
	filenames stringChan
	events    elementChan
}

func (wc *workerPool) startWorkers(n int) {
	for w := 1; w <= n; w++ {
		wc.waiter.Add(1)
		go wc.worker(w)
	}

	go func() {
		wc.waiter.Wait()
		close(wc.events)
		log.Println("Events EOD")
	}()
}

func (wc *workerPool) populateJobs(patterns []string) {
	defer close(wc.filenames)

	for k, pattern := range patterns {
		log.Printf("[IN] - %d -", k)

		files, _ := filepath.Glob(pattern)
		for _, file := range files {
			wc.filenames <- file
			log.Printf("[IN] %q", file)
		}
	}

	log.Println("[IN] DONE")
}

func EventChannel(searchPaths []string, workerPoolSize int) <-chan *dom.Element {
	wc := newWorkerPool()
	wc.startWorkers(workerPoolSize)
	go wc.populateJobs(searchPaths)

	return wc.events
}

const eventName = "ZDARZENIE"

func (wc *workerPool) worker(n int) {
	defer wc.waiter.Done()

	log.Printf("[WORKER] %d START", n)

	for filename := range wc.filenames {
		log.Printf("[WORKER] %d [BEGINS] %q\n%s", n, filename, sys.MemStats())

		doc, err := parse(filename)
		if err != nil {
			log.Println(filename + ": " + err.Error())

			continue
		}

		for _, e := range diveTo(eventName, doc.Root.Children) {
			wc.events <- e
		}

		log.Printf("[WORKER] %d [FINISHED] %q\n%s", n, filename, sys.MemStats())
	}

	log.Printf("[WORKER] %d STOP", n)
}

func diveTo(s string, children []*dom.Element) []*dom.Element {
	for _, e := range children {
		if e.Name == s {
			return children
		}

		return diveTo(s, e.Children)
	}

	log.Printf(`cannot find "%s" elements`, s)

	return nil
}

func parse(filename string) (*dom.Document, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return xml.Parse(file)
}
