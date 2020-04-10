package sewik

import (
	"fmt"
	"log"
	"os"
	"sync/atomic"

	"github.com/subchen/go-xmldom"

	"sewik/pkg/es"
	"sewik/pkg/sync"
	"sewik/pkg/xml"
)

func ElasticDocs(n <-chan *xmldom.Node) <-chan *es.Doc {
	documents := make(chan *es.Doc, cap(n))

	for nn := range n {
		documents <- es.NewDoc(nn)
	}
	close(documents)

	return documents
}

func ElementsOf(elementName string, filenames <-chan string, workerLimit int, size int) <-chan *xmldom.Node {
	wg := sync.LimitingWaitGroup{Limit: workerLimit + 1}

	elements := make(chan *xmldom.Node, size)
	go func() {
		wg.Wait()
		close(elements)
	}()

	var n uint32

	go func() {
		wg.Add(1)
		defer wg.Done()
		for filename := range filenames {
			wg.Add(1)
			go func(filename string) {
				defer wg.Done()
				atomic.AddUint32(&n, 1)
				log.Printf("[STARTED] %d %q", n, filename)

				doc, err := parse(filename)
				if err != nil {
					log.Println(err)
					return
				}

				for _, e := range dive(elementName, doc.Root.Children) {
					e.SetAttributeValue("_src", filename)
					elements <- e
				}

				log.Printf("[FINISHED] %d %q", n, filename)
			}(filename)
		}
	}()

	return elements
}

func parse(filename string) (*xmldom.Document, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("%q: %s", filename, err)
	}
	defer file.Close()

	return xml.Parse(file)
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
