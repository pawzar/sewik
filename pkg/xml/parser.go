package xml

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"

	"github.com/subchen/go-xmldom"
	"golang.org/x/net/html/charset"
)

func Parse(r io.Reader) (*xmldom.Document, error) {
	p := xml.NewDecoder(r)
	p.CharsetReader = charset.NewReaderLabel

	t, err := p.Token()
	if err != nil {
		return nil, err
	}

	doc := &xmldom.Document{}
	var e *xmldom.Node
	for t != nil {
		switch token := t.(type) {
		case xml.StartElement:
			el := &xmldom.Node{}
			el.Document = doc
			el.Parent = e
			el.Name = token.Name.Local
			for _, attr := range token.Attr {
				el.Attributes = append(el.Attributes, &xmldom.Attribute{
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
			if e != nil {
				e.Text = string(bytes.TrimSpace(token))
			}
		case xml.ProcInst:
			doc.ProcInst = stringifyProcInst(&token)
		case xml.Directive:
			doc.Directives = append(doc.Directives, stringifyDirective(&token))
		}

		t, err = p.Token()
	}

	if err != io.EOF {
		return nil, err
	}

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
