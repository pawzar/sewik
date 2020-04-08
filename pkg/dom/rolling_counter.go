package dom

import (
	json2 "encoding/json"
	"log"
	"sync"
)

func NewRollingCounter() *RollingCounter {
	return &RollingCounter{
		F: make(rollingCounterFields),
	}
}

type rollingCounterFields map[string]*RollingCounter

type RollingCounter struct {
	m sync.RWMutex
	C int
	F rollingCounterFields
}

func (c *RollingCounter) MarshalJSON() ([]byte, error) {
	if len(c.F) == 1 {
		for _, v := range c.F {
			if v.C == 1 {
				for _, x := range c.F {
					return json2.Marshal(x)
				}
			}
			return json2.Marshal(v)
		}
	}
	return json2.Marshal(c.F)
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
	if i > c.C {
		c.C = i
	}
}

func (c *RollingCounter) addToChild(s string, v *Counter) {
	c.find(s).Add(v)
}

func (c *RollingCounter) find(s string) *RollingCounter {
	c.m.Lock()
	defer c.m.Unlock()
	ff, ok := c.F[s]
	if !ok {
		ff = NewRollingCounter()
		c.F[s] = ff
	}
	return ff
}
