package dom

import (
	"fmt"
	"strings"

	"github.com/subchen/go-xmldom"
)

type Info struct {
	counters map[string]int
	children map[string]*Info
}

func (i Info) GoString() string {
	if len(i.children) > 0 {
		var s strings.Builder
		s.WriteString(fmt.Sprintf("Info{\ncounters: %#v,\nchildren: map[string]*Info{\n", i.counters))
		for k, v := range i.children {
			s.WriteString(fmt.Sprintf("%q:&%#v,\n", k, v))
		}
		s.WriteString("},\n}")
		return s.String()
	}
	return fmt.Sprintf("Info{counters: %#v}", i.counters)
}

func NewInfo() *Info {
	return &Info{
		counters: make(map[string]int),
		children: make(map[string]*Info),
	}
}

func newInfoFrom(n *xmldom.Node) *Info {
	i := NewInfo()
	i.count(n)
	return i
}

func (i *Info) count(n *xmldom.Node) {
	if n.Text != "" {
		i.counters["_"]++
	}
	for _, v := range n.Attributes {
		i.counters["_"+v.Name]++
	}
	for _, v := range n.Children {
		i.counters[v.Name]++
	}
}

func (i *Info) String() string {
	return i.string("") + "\n"
}

func (i *Info) string(p string) string {
	var s strings.Builder
	for k, v := range i.counters {
		s.WriteString(fmt.Sprintf("%s%s = %d\n", p, k, v))
		if x, ok := i.children[k]; ok {
			s.WriteString(x.string("\t" + p))
		}
	}
	return s.String()
}

func (i *Info) Add(n *xmldom.Node) {
	i.apply(newInfoFrom(n))
	for _, v := range n.Children {
		_, ok := i.children[v.Name]
		if !ok {
			i.children[v.Name] = NewInfo()
		}
		i.children[v.Name].Add(v)
	}
}

func (i *Info) apply(in *Info) {
	for k, v := range in.counters {
		x := i.counters[k]
		if v > x {
			i.counters[k] = v
		}
	}
}
