package dom

import (
	json2 "encoding/json"
	"log"
	"sync"
)

func NewRollingCounter() *RollingCounter {
	return &RollingCounter{
		Fields: make(rollingCounterFields),
	}
}

type rollingCounterFields map[string]*RollingCounter

type RollingCounter struct {
	m      sync.RWMutex
	Count  int
	Fields rollingCounterFields
}

func (c *RollingCounter) String() string {
	bytes, err := json2.Marshal(c)
	if err != nil {
		log.Panic(err)
	}
	return string(bytes)
}

func (c *RollingCounter) Add(n *Counter) {
	c.add(n.c)
	for k, v := range n.f {
		c.addToChild(k, v)
	}
}

func (c *RollingCounter) add(i int) {
	c.m.Lock()
	defer c.m.Unlock()
	if i > c.Count {
		c.Count = i
	}
}

func (c *RollingCounter) addToChild(s string, v *Counter) {
	c.find(s).Add(v)
}

func (c *RollingCounter) find(s string) *RollingCounter {
	c.m.Lock()
	defer c.m.Unlock()
	ff, ok := c.Fields[s]
	if !ok {
		ff = NewRollingCounter()
		c.Fields[s] = ff
	}
	return ff
}
