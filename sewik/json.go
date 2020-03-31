package sewik

import (
	"fmt"
	"io"
	"strings"

	"sewik/xml/dom"
)

const IdElementName = "ID"

type JSON struct {
	Id   string
	Body io.Reader
}

func (o JSON) Read(p []byte) (n int, err error) {
	return o.Body.Read(p)
}

func NewJSON(element *dom.Element) (event JSON) {
	for _, c := range element.Children {
		if c.Name == IdElementName {
			event.Id = c.Text
		}
	}

	var b strings.Builder
	json(&b, element)

	event.Body = strings.NewReader(b.String())

	return
}

func addAttr(dest *strings.Builder, a *dom.Attribute) {
	if dest.Len() > 0 {
		fmt.Fprint(dest, `,`)
	}
	fmt.Fprintf(dest, `_%q:"%q"`, a.Name, a.Value)
}

func addElem(dest *strings.Builder, e *dom.Element) {
	if dest.Len() > 0 {
		fmt.Fprint(dest, `,`)
	}

	fmt.Fprintf(dest, `"%q":`, e.Name)

	json(dest, e)
}

func json(dest *strings.Builder, e *dom.Element) {
	if len(e.Attributes)+len(e.Children) == 0 {
		jsValue(dest, e)

		return
	}

	jsObject(dest, e)
}

func jsValue(dest *strings.Builder, e *dom.Element) (int, error) {
	return fmt.Fprintf(dest, `"%q"`, e.Text)
}

func jsObject(dest *strings.Builder, e *dom.Element) {
	var s strings.Builder

	if e.Text != "" {
		fmt.Fprintf(&s, `"_"="%q"`, e.Text)
	}

	for _, a := range e.Attributes {
		addAttr(&s, a)
	}

	for _, c := range e.Children {
		addElem(&s, c)
	}

	fmt.Fprintf(dest, `{%s}`, s)
}
