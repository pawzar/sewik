package sewik

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"sewik/pkg/dom"
	"sewik/pkg/es"
	"sewik/pkg/sync"
	"sewik/pkg/sys"
	"sewik/pkg/xml"
)

func Docs(elementName string, filenames <-chan string, workerLimit int, size int) <-chan *es.Document {
	wg := sync.LimitingWaitGroup{Limit: workerLimit}

	documents := make(chan *es.Document, size)
	sys.ChUtDo("doc", documents)
	go func() {
		wg.Wait()
		close(documents)
	}()

	go func() {
		n := 1
		for filename := range filenames {
			wg.Add(1)
			go func(n int, filename string) {
				defer wg.Done()
				log.Printf("[STARTED] %d %q", n, filename)

				doc, err := parse(filename)
				if err != nil {
					log.Println(err)
					return
				}

				for _, e := range dive(elementName, doc.Root.Children) {
					documents <- es.NewDoc(e, filename)
				}

				log.Printf("[FINISHED] %d %q", n, filename)
			}(n, filename)
			n++
		}
	}()

	return documents
}
func ElementsOf(elementName string, filenames <-chan string, workerLimit int) <-chan *dom.Element {
	wg := sync.LimitingWaitGroup{Limit: workerLimit}

	elements := make(chan *dom.Element, 1000)
	sys.ChUtEl("els", elements)
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
					e.SetAttributeValue("_src", strings.Replace(strconv.QuoteToASCII(filename), `"`, "", -1))
					elements <- e
				}

				log.Printf("[FINISHED] %d %q", n, filename)
			}(filename)
		}
		n++
	}()

	return elements
}

func RootElements(workerLimit int, filenames <-chan string) <-chan *dom.Element {
	wg := sync.LimitingWaitGroup{Limit: workerLimit}

	roots := make(chan *dom.Element)
	sys.ChUtEl("rts", roots)
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

func parse(filename string) (*dom.Document, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("%q: %s", filename, err)
	}
	defer file.Close()

	return xml.Parse(file)
}

func dive(s string, children []*dom.Element) []*dom.Element {
	for _, e := range children {
		if e.Name == s {
			return children
		}

		return dive(s, e.Children)
	}

	log.Printf(`cannot find "%s" elements`, s)

	return nil
}
