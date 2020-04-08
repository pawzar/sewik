package dom

import (
	json2 "encoding/json"
	"log"

	"github.com/subchen/go-xmldom"
)

type Object struct {
	Value  string
	Fields fields
}

func (o *Object) MarshalJSON() ([]byte, error) {
	return json2.Marshal(&o.Fields)
}

type fields map[string]*Object

func NewObject() *Object {
	return &Object{
		Fields: make(fields),
	}
}
func (o *Object) From(n *xmldom.Node) *Object {
	o.Add(n)
	return o
}
func (o *Object) Add(n *xmldom.Node) {
	o.addNode(n.Children...)
	o.addAttr(n.Attributes...)
	o.addText(n.Text)
}

func (o *Object) String() string {
	bytes, err := json2.Marshal(o)
	if err != nil {
		log.Panic(err)
	}
	return string(bytes)
}

func (o *Object) add(s, v string) {
	o.find(s).Value = v
}

func (o *Object) addText(t string) {
	if t == "" {
		return
	}

	if len(o.Fields) > 0 {
		o.add("_", t)
	}
	o.Value = t
}

func (o *Object) addAttr(a ...*xmldom.Attribute) {
	for _, aa := range a {
		o.add("_"+aa.Name, aa.Value)
	}
}

func (o *Object) addNode(n ...*xmldom.Node) {
	for _, nn := range n {
		o.find(nn.Name).Add(nn)
	}
}

func (o *Object) find(s string) *Object {
	ff, ok := o.Fields[s]
	if !ok {
		ff = NewObject()
		o.Fields[s] = ff
	}
	return ff
}
