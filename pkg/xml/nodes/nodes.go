package nodes

import (
	"github.com/subchen/go-xmldom"

	"sewik/pkg/json"
)

func NewMap(node *xmldom.Node) json.Map {
	return json.NewMapAs(node.Name, NewInnerMap(node))
}

func NewInnerMap(node *xmldom.Node) json.Map {
	mm := json.NewMap()
	mm.AddStringAs("_", node.Text)
	for _, a := range node.Attributes {
		mm.AddStringAs("_"+a.Name, a.Value)
	}
	for _, c := range node.Children {
		mm.AddFrom(NewMap(c))
	}
	return mm
}
