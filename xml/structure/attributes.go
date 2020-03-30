package structure

import (
	"sync"

	"github.com/subchen/go-xmldom"
)

type Attributes interface {
	Add(n *xmldom.Attribute)
	Int() attributesMap
	Len() int
}

func NewAttributesWithLock() Attributes {
	return &attributesWithLock{
		in: make(attributesMap),
	}
}

type attributesMap map[string]int

type attributesWithLock struct {
	mx sync.Mutex
	in attributesMap
}

func (a *attributesWithLock) Add(n *xmldom.Attribute) {
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

func (a attributesWithLock) Int() attributesMap {
	return a.in
}

func (a attributesWithLock) Len() int {
	return len(a.in)
}
