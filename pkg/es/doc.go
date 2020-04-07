package es

import (
	json2 "encoding/json"
	"log"

	"github.com/subchen/go-xmldom"

	"sewik/pkg/json"
	"sewik/pkg/xml/nodes"
)

type Doc struct {
	ID     string
	Source string
	m      json.Map
}

func (d Doc) String() string {
	bytes, err := json2.Marshal(d.m)
	if err != nil {
		log.Panic(err)
	}
	return string(bytes)
}

func NewDoc(node *xmldom.Node) *Doc {
	return &Doc{
		ID:     findID(node),
		Source: findSrc(node),
		m:      nodes.NewInnerMap(node),
	}
}
func NewDocWithSrc(node *xmldom.Node, src string) *Doc {
	return &Doc{
		ID:     findID(node),
		Source: src,
		m:      nodes.NewMap(node),
	}
}

func findSrc(node *xmldom.Node) string {
	var src string
	for _, a := range node.Attributes {
		if a.Name == "_src" {
			src = a.Value
		}
	}
	return src
}
func findID(node *xmldom.Node) string {
	var id string
	for _, c := range node.Children {
		if c.Name == "ID" {
			id = c.Text
		}
	}
	return id
}

var _ = []string{
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
