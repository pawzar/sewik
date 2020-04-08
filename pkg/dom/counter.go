package dom

import (
	json2 "encoding/json"
	"log"
	"sync"

	"github.com/subchen/go-xmldom"
)

func NewCounter() *Counter {
	return &Counter{
		f: make(counterFields),
	}
}

type Counter struct {
	m sync.RWMutex
	c int
	f counterFields
}

type counterFields map[string]*Counter

func (c *Counter) String() string {
	bytes, err := json2.Marshal(c)
	if err != nil {
		log.Panic(err)
	}
	return string(bytes)
}

func (c *Counter) WithNode(n *xmldom.Node) *Counter {
	c.Add(n)
	return c
}

func (c *Counter) Add(n *xmldom.Node) {
	c.add()
	c.addFromText(n.Text)
	c.addFromAttr(n.Attributes...)
	c.addFromNode(n.Children...)
}

func (c *Counter) add() {
	c.m.Lock()
	defer c.m.Unlock()
	c.c++
}

func (c *Counter) addFromText(t string) {
	if t == "" {
		return
	}
	c.count("_")
}

func (c *Counter) addFromAttr(a ...*xmldom.Attribute) {
	for _, aa := range a {
		c.count("_" + aa.Name)
	}
}

func (c *Counter) addFromNode(n ...*xmldom.Node) {
	for _, nn := range n {
		c.find(nn.Name).Add(nn)
	}
}

func (c *Counter) count(s string) {
	c.find(s).add()
}

func (c *Counter) find(s string) *Counter {
	c.m.Lock()
	defer c.m.Unlock()
	ff, ok := c.f[s]
	if !ok {
		ff = NewCounter()
		c.f[s] = ff
	}
	return ff
}
