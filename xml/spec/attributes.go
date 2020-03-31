package spec

import (
	"sync"

	"sewik/xml/dom"
)

type Attributes interface {
	Add(n *dom.Attribute)
	Get() attributesMap
	Len() int
}
type Attribute = int

func NewAttributesWithLock() Attributes {
	return &attributesWithLock{
		in: make(attributesMap),
	}
}

type attributesMap map[string]Attribute

type attributesWithLock struct {
	mx sync.Mutex
	in attributesMap
}

func (a *attributesWithLock) Add(n *dom.Attribute) {
	a.mx.Lock()
	defer a.mx.Unlock()

	x, exists := a.in[n.Name]

	if exists {
		x++
	} else {
		x = 1
	}

	a.in[n.Name] = x
}

func (a attributesWithLock) Get() attributesMap {
	return a.in
}

func (a attributesWithLock) Len() int {
	return len(a.in)
}
