package xml

import (
	"fmt"
	"log"
	"os"
	"sync/atomic"

	"github.com/subchen/go-xmldom"

	"sewik/internal/sync"
)

func ElementsOf(elementName string, filenames <-chan string, workerLimit int, size int) <-chan *xmldom.Node {
	wg := sync.SemaphoredWaitGroup{Size: workerLimit + 1}

	elements := prepareChannel(size, &wg)

	go func() {
		wg.Add(1)
		defer wg.Done()
		var n uint32
		for filename := range filenames {
			wg.Add(1)
			go func(filename string) {
				defer wg.Done()
				nn := atomic.AddUint32(&n, 1)

				log.Printf("[STARTED] %d %q", nn, filename)

				doc, err := parse(filename)
				if err != nil {
					log.Println(err)
					return
				}

				for _, e := range dive(elementName, doc.Root.Children) {
					e.SetAttributeValue("_src", filename)
					elements <- e
				}

				log.Printf("[FINISHED] %d %q", nn, filename)
			}(filename)
		}
	}()

	return elements
}

func prepareChannel(size int, wg sync.WaitGroup) chan *xmldom.Node {
	elements := make(chan *xmldom.Node, size)
	go func() {
		wg.Wait()
		close(elements)
	}()
	return elements
}

func parse(filename string) (*xmldom.Document, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("%q: %s", filename, err)
	}
	defer file.Close()

	return Parse(file)
}

func dive(s string, children []*xmldom.Node) []*xmldom.Node {
	for _, e := range children {
		if e.Name == s {
			return children
		}

		return dive(s, e.Children)
	}

	log.Printf(`cannot find "%s" elements`, s)

	return nil
}
