package json

type Map map[string][]interface{}

func NewMap() Map {
	return make(Map)
}
func NewMapFrom(mm Map) Map {
	n := NewMap()
	n.AddFrom(mm)
	return n
}
func NewMapAs(name string, mm Map) Map {
	n := NewMap()
	n.AddAs(name, mm)
	return n
}
func (m Map) AddAs(key string, mm Map) {
	m[key] = append(m[key], mm)
}
func (m Map) AddFrom(mm Map) {
	for key, items := range mm {
		m[key] = append(m[key], items...)
	}
}
func (m Map) AddStringAs(key string, vv ...string) {
	for _, v := range vv {
		if v != "" {
			m[key] = append(m[key], v)
		}
	}
}
