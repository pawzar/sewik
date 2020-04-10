package mapping

import (
	"github.com/subchen/go-xmldom"

	"sewik/pkg/dom"
)

type Mapping struct {
	i *dom.Info
	n *xmldom.Node
}

func NewMapping(i *dom.Info, n *xmldom.Node) *Mapping {
	return &Mapping{i: i, n: n}
}

func (i *Mapping) Map() interface{} {
	if i.i.IsObsolete() {
		for _, v := range i.n.Children {
			return NewMapping(i.i.Get(v.Name), v).Map()
		}
		for _, v := range i.n.Attributes {
			return v.Value
		}
		return i.n.Text
	}

	m := make(map[string]interface{})
	a := make(map[string][]interface{})

	for _, v := range i.n.Children {
		if i.i.IsArray(v.Name) {
			a[v.Name] = append(a[v.Name], NewMapping(i.i.Get(v.Name), v).Map())
		} else {
			m[v.Name] = NewMapping(i.i.Get(v.Name), v).Map()
		}
	}
	for _, v := range i.n.Attributes {
		if i.i.IsArray(v.Name) {
			a[v.Name] = append(a[v.Name], v.Value)
		} else {
			m[v.Name] = v.Value
		}
	}

	for k, v := range a {
		m[k] = v
	}

	return m
}
