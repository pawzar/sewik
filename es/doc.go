package es

import (
	"fmt"
	"io"
	"io/ioutil"
	"strings"

	"sewik/dom"
)

const IdElementName = "ID"

var x = []string{
	"UCZESTNICY",
	"POJAZDY",
}

func isArray(name string) bool {
	for _, n := range x {
		if n == name {
			return true
		}
	}
	return false
}

type Document struct {
	Id   string
	Body io.Reader
}

func (o Document) String() string {
	body, _ := ioutil.ReadAll(o.Body)

	return fmt.Sprintf(`"%s":%s`, o.Id, string(body))
}

func NewDoc(element *dom.Element) *Document {
	event := Document{}

	for _, c := range element.Children {
		if c.Name == IdElementName {
			event.Id = c.Text
		}
	}

	var b strings.Builder
	json(&b, element)

	event.Body = strings.NewReader(b.String())

	return &event
}

func json(dest *strings.Builder, e *dom.Element) {
	if isArray(e.Name) {
		jsArray(dest, e)

		return
	}

	if len(e.Attributes)+len(e.Children) == 0 {
		jsValue(dest, e)

		return
	}

	jsObject(dest, e)
}

func jsValue(dest *strings.Builder, e *dom.Element) {
	fmt.Fprintf(dest, `"%s"`, e.Text)
}

func jsArray(dest *strings.Builder, e *dom.Element) {
	var s strings.Builder

	for _, c := range e.Children {
		addElem(&s, c, true)
	}

	fmt.Fprintf(dest, `[%s]`, s.String())
}

func jsObject(dest *strings.Builder, e *dom.Element) {
	var s strings.Builder

	if e.Text != "" {
		fmt.Fprintf(&s, `"__text":"%s"`, e.Text)
	}

	for _, a := range e.Attributes {
		addAttr(&s, a)
	}

	for _, c := range e.Children {
		addElem(&s, c, false)
	}

	fmt.Fprintf(dest, `{%s}`, s.String())
}

func addAttr(dest *strings.Builder, a *dom.Attribute) {
	if dest.Len() > 0 {
		fmt.Fprint(dest, `,`)
	}

	fmt.Fprintf(dest, `"_%s":"%s"`, a.Name, a.Value)
}

func addElem(dest *strings.Builder, e *dom.Element, internalizeName bool) {
	if dest.Len() > 0 {
		fmt.Fprint(dest, `,`)
	}

	if internalizeName {
		e.SetAttributeValue("_type", e.Name)
	} else {
		fmt.Fprintf(dest, `"%s":`, e.Name)
	}

	json(dest, e)
}
