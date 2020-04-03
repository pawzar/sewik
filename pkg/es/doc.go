package es

import (
	"fmt"
	"strconv"
	"strings"

	"sewik/pkg/dom"
	"sewik/pkg/json"
)

const IdElementName = "ID"

var x = []string{
	"UCZESTNICY",
	"POJAZDY",
	"NAWIERZCHNIA",
	"STAN_NAWIERZCHNI",
	"RODZAJ_DROGI",
	"SYGNALIZACJA_SWIETLNA",
	"OZNAKOWANIE_POZIOME",
	"OBSZAR_ZABUDOWANY",
	"CHARAKT_MIEJSCA",
	"WARUNKI_ATMOSFERYCZNE",
	"USZKODZENIA_POZA_POJAZDAMI",
	"INNE_PRZYCZYNY",
	"GEOMETRIA_DROGI",
	"INFO_O_DRODZE",
	"STAN_POJAZDU",
	"MIEJSCE",
	"INNE_CECHY_POJAZU",
	"INFO_DODATKOWE",
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
	ID     string
	Body   string
	Source string
	fs     string
}

func (o Document) String() string {
	return fmt.Sprintf(`{"id":"%s","src":"%s","body":`+o.fs+`}`, o.ID, o.Source, o.Body)
}

func NewDoc(element *dom.Element, src string) *Document {
	o := Document{
		Source: json.Escape(strings.Trim(strconv.QuoteToASCII(src), `"`)),
		fs:     "%s",
	}

	for _, c := range element.Children {
		if c.Name == IdElementName {
			o.ID = json.Escape(c.Text)
		}
	}

	var b strings.Builder
	o.fs = jsonize(&b, element, false)
	o.Body = b.String()

	return &o
}

func jsonize(dest *strings.Builder, e *dom.Element, wrap bool) string {
	if isArray(e.Name) {
		return jsArray(dest, e, wrap)
	}

	if len(e.Attributes)+len(e.Children) == 0 {
		return jsValue(dest, e, wrap)
	}

	return jsObject(dest, e, wrap)
}

func jsValue(dest *strings.Builder, e *dom.Element, wrap bool) string {
	fs := `"%s"`
	if !wrap {
		fmt.Fprintf(dest, `%s`, json.Escape(e.Text))
	} else {
		fmt.Fprintf(dest, fs, json.Escape(e.Text))
	}
	return fs
}

func jsArray(dest *strings.Builder, e *dom.Element, wrap bool) string {
	var s strings.Builder
	fs := `[%s]`

	for _, c := range e.Children {
		addElem(&s, c, true)
	}

	if !wrap {
		fmt.Fprintf(dest, `%s`, s.String())
	} else {
		fmt.Fprintf(dest, fs, s.String())
	}
	return fs
}

func jsObject(dest *strings.Builder, e *dom.Element, wrap bool) string {
	var s strings.Builder
	fs := `{%s}`

	if e.Text != "" {
		fmt.Fprintf(&s, `"__text":"%s"`, json.Escape(e.Text))
	}

	for _, a := range e.Attributes {
		addAttr(&s, a)
	}

	for _, c := range e.Children {
		addElem(&s, c, false)
	}

	if !wrap {
		fmt.Fprintf(dest, `%s`, s.String())
	} else {
		fmt.Fprintf(dest, fs, s.String())
	}
	return fs
}

func addAttr(dest *strings.Builder, a *dom.Attribute) {
	if dest.Len() > 0 {
		fmt.Fprint(dest, `,`)
	}

	fmt.Fprintf(dest, `"_%s":"%s"`, json.Escape(a.Name), json.Escape(a.Value))
}

func addElem(dest *strings.Builder, e *dom.Element, internalizeName bool) {
	if dest.Len() > 0 {
		fmt.Fprint(dest, `,`)
	}

	if internalizeName {
		e.SetAttributeValue("_type", json.Escape(e.Name))
	} else {
		fmt.Fprintf(dest, `"%s":`, json.Escape(e.Name))
	}

	jsonize(dest, e, true)
}
