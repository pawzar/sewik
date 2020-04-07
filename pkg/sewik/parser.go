package sewik

import (
	"fmt"
	"log"
	"os"

	"github.com/subchen/go-xmldom"

	"sewik/pkg/es"
	"sewik/pkg/sync"
	"sewik/pkg/xml"
)

func EsDocs(in <-chan *xmldom.Node) <-chan *es.Doc {
	documents := make(chan *es.Doc, cap(in))

	for e := range in {
		documents <- es.NewDoc(e)
	}
	close(documents)

	return documents
}

func ElementsOf(elementName string, filenames <-chan string, workerLimit int, size int) <-chan *xmldom.Node {
	wg := sync.LimitingWaitGroup{Limit: workerLimit}

	elements := make(chan *xmldom.Node, size)
	go func() {
		wg.Wait()
		close(elements)
	}()

	go func() {
		n := 1
		for filename := range filenames {
			wg.Add(1)
			go func(filename string) {
				defer wg.Done()
				log.Printf("[STARTED] %d %q", n, filename)

				doc, err := parse(filename)
				if err != nil {
					log.Println(err)
					return
				}

				for _, e := range dive(elementName, doc.Root.Children) {
					e.SetAttributeValue("src", filename)
					elements <- e
				}

				log.Printf("[FINISHED] %d %q", n, filename)
			}(filename)
		}
		n++
	}()

	return elements
}

func Roots(filenames <-chan string, workerLimit int, size int) <-chan *xmldom.Node {
	wg := sync.LimitingWaitGroup{Limit: workerLimit}

	roots := make(chan *xmldom.Node, size)
	go func() {
		wg.Wait()
		close(roots)
	}()

	go func() {
		n := 1
		for filename := range filenames {
			wg.Add(1)
			log.Printf("[DISPATCH] %d %q", n, filename)

			go func(n int, filename string) {
				defer wg.Done()
				log.Printf("[START] %d %q", n, filename)

				fmt.Printf("<!-- %s -->\n", filename)

				doc, err := parse(filename)
				if err != nil {
					log.Println(err)
					return
				}

				roots <- doc.Root

				log.Printf("[STOP] %d %q", n, filename)
			}(n, filename)

			n++
		}
	}()

	return roots
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
