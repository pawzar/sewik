package mapping

import (
	"github.com/subchen/go-xmldom"

	"sewik/internal/dom"
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
		name := v.Name
		if i.i.IsArray(name) {
			a[name] = append(a[name], NewMapping(i.i.Get(name), v).Map())
		} else {
			m[name] = NewMapping(i.i.Get(name), v).Map()
		}
	}

	for _, v := range i.n.Attributes {
		name := "_" + v.Name
		if i.i.IsArray(name) {
			a[name] = append(a[name], v.Value)
		} else {
			m[name] = v.Value
		}
	}

	for k, v := range a {
		m[k] = v
	}

	return m
}
