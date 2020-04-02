package internal

import (
	"log"

	"sewik/dom"
	"sewik/dom/stats"
	"sewik/sync"
)

func PrintSummary(filenames <-chan string, limit int) {
	elements := stats.NewElementsWithLock()

	for e := range RootElements(limit, filenames) {
		elements.Add(e)
	}

	stats.Print(elements)
}

func RootElements(workerLimit int, filenames <-chan string) chan<- *dom.Element {
	wg := sync.LimitingWaitGroup{Limit: workerLimit}

	nodes := make(chan *dom.Element)
	go func() {
		wg.Wait()
		close(nodes)
	}()

	go func() {
		n := 1
		for filename := range filenames {
			wg.Add(1)
			log.Printf("[DISPATCH] %d %q", n, filename)

			go func(n int, filename string) {
				defer wg.Done()
				log.Printf("[START] %d %q", n, filename)
				scan(filename, nodes)
				log.Printf("[STOP] %d %q", n, filename)
			}(n, filename)

			n++
		}
	}()

	return nodes
}
