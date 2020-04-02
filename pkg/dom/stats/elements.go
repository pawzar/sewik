package stats

import (
	"sync"

	"sewik/pkg/dom"
)

type Elements interface {
	Add(n *dom.Element)
	Get() elementMap
	Len() int
}

type Element struct {
	Cn int
	At Attributes
	El Elements
}

func NewElementsWithLock() Elements {
	return &elementsWithLock{
		in: make(elementMap),
	}
}

type elementsWithLock struct {
	mx sync.Mutex
	in elementMap
}

type elementMap map[string]Element

func (e *elementsWithLock) Add(n *dom.Element) {
	e.mx.Lock()
	defer e.mx.Unlock()

	x, exists := e.in[n.Name]

	if exists {
		x.Cn++
	} else {
		x = Element{
			Cn: 1,
			At: newAttributesWithLock(),
			El: NewElementsWithLock(),
		}
	}

	for _, a := range n.Attributes {
		x.At.Add(a)
	}
	for _, c := range n.Children {
		x.El.Add(c)
	}

	e.in[n.Name] = x
}

func (e elementsWithLock) Get() elementMap {
	return e.in
}

func (e elementsWithLock) Len() int {
	return len(e.in)
}