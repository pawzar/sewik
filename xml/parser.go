package xml

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"

	"golang.org/x/net/html/charset"

	"sewik/dom"
)

func Parse(r io.Reader) (*dom.Document, error) {
	p := xml.NewDecoder(r)
	p.CharsetReader = charset.NewReaderLabel

	t, err := p.Token()
	if err != nil {
		return nil, err
	}

	doc := &dom.Document{}
	var e *dom.Element
	for t != nil {
		switch token := t.(type) {
		case xml.StartElement:
			// a new node
			el := &dom.Element{}
			el.Document = doc
			el.Parent = e
			el.Name = token.Name.Local
			for _, attr := range token.Attr {
				el.Attributes = append(el.Attributes, &dom.Attribute{
					Name:  attr.Name.Local,
					Value: attr.Value,
				})
			}
			if e != nil {
				e.Children = append(e.Children, el)
			}
			e = el

			if doc.Root == nil {
				doc.Root = e
			}
		case xml.EndElement:
			e = e.Parent
		case xml.CharData:
			// text node
			if e != nil {
				e.Text = string(bytes.TrimSpace(token))
			}
		case xml.ProcInst:
			doc.ProcInst = stringifyProcInst(&token)
		case xml.Directive:
			doc.Directives = append(doc.Directives, stringifyDirective(&token))
		}

		// get the next token
		t, err = p.Token()
	}

	// Make sure that reading stopped on EOF
	if err != io.EOF {
		return nil, err
	}

	// All is good, return the document
	return doc, nil
}

func stringifyProcInst(pi *xml.ProcInst) string {
	if pi == nil {
		return ""
	}
	return fmt.Sprintf("<?%s %s?>", pi.Target, string(pi.Inst))
}

func stringifyDirective(directive *xml.Directive) string {
	if directive == nil {
		return ""
	}
	return fmt.Sprintf("<!%s>", string(*directive))
}
