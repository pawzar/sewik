package docs

import (
	"github.com/subchen/go-xmldom"

	"sewik/pkg/json"
	"sewik/pkg/xml/nodes"
)

func NewMap(d *xmldom.Document) json.Map {
	mm := json.NewMapFrom(nodes.NewMap(d.Root))
	mm.AddStringAs("__proc", d.ProcInst)
	mm.AddStringAs("_directives", d.Directives...)
	return mm
}
