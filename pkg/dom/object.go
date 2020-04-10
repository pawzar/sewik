package dom

import (
	"encoding/json"
	"log"

	"github.com/subchen/go-xmldom"
)

type Object struct {
	Text     string
	Fields   map[string]string
	Children map[string][]*Object
}

func (o *Object) MarshalJSON() ([]byte, error) {
	if o.Text != "" {
		return json.Marshal(o.Text)
	}
	if len(o.Children) == 1 {
		for _, v := range o.Children {
			if len(v) == 1 {
				for _, w := range v {
					return json.Marshal(w)
				}
			}
			return json.Marshal(v)
		}
	}
	return json.Marshal(o.Children)
}

func NewObject() *Object {
	return &Object{
		Fields:   make(map[string]string),
		Children: make(map[string][]*Object),
	}
}

func (o *Object) With(n *xmldom.Node) *Object {
	o.Add(n)
	return o
}

func (o *Object) WithText(s string) *Object {
	o.addText(s)
	return o
}

func (o *Object) Add(n *xmldom.Node) {
	o.addText(n.Text)
	o.addAttr(n.Attributes...)
	o.addNode(n.Children...)
}

func (o *Object) String() string {
	bytes, err := json.Marshal(o)
	if err != nil {
		log.Panic(err)
	}
	return string(bytes)
}

func (o *Object) addText(t string) {
	if t == "" {
		return
	}
	o.Text = t
}

func (o *Object) addAttr(a ...*xmldom.Attribute) {
	for _, aa := range a {
		o.Fields[aa.Name] = aa.Value
	}
}

func (o *Object) addNode(n ...*xmldom.Node) {
	for _, nn := range n {
		if nn.Text != "" {
			o.Fields[nn.Name] = nn.Text
		} else {
			o.append(nn.Name, NewObject().With(nn))
		}
	}
}

func (o *Object) append(s string, oo *Object) {
	ff, _ := o.Children[s]
	objects := append(ff, oo)
	o.Children[s] = objects
}
