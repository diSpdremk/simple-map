package simple_map

import "testing"

func TestMap(t *testing.T) {
	m := NewSimpleMap[int, int]()
	m.Set(1, 1)
	m.Delete(1)
}
