package stats

import (
	"sync"

	"github.com/subchen/go-xmldom"
)

type Attributes interface {
	Add(n *xmldom.Attribute)
	Get() attributeMap
	Len() int
}
type attribute = int

func newAttributesWithLock() Attributes {
	return &attributesWithLock{
		in: make(attributeMap),
	}
}

type attributeMap map[string]attribute

type attributesWithLock struct {
	mx sync.Mutex
	in attributeMap
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

func (a attributesWithLock) Get() attributeMap {
	return a.in
}

func (a attributesWithLock) Len() int {
	return len(a.in)
}
