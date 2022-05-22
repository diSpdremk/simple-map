package simple_map

import (
	"encoding/json"
	"github.com/zhenjl/cityhash"
	"math"
)

type element[KT, VT any] struct {
	k KT
	v VT
	n *element[KT, VT]
}

func (s *SMap[KT, VT]) unsafeSet(index int, k KT, v VT) {
	e := s.Bucket[index]
	if e == nil {
		s.Bucket[index] = &element[KT, VT]{
			k: k,
			v: v,
		}
		return
	}
	for e.k != k && e.n != nil {
		e = e.n
	}
	if e.k == k {
		e.v = v
		return
	}
	e.n = &element[KT, VT]{
		k: k,
		v: v,
	}
}

func (s *SMap[KT, VT]) unsafeGet(index int, k KT) (VT, bool) {
	var dV VT
	e := s.Bucket[index]
	for e.k != k {
		e = e.n
	}
	if e == nil {
		return dV, false
	}
	return e.v, true
}

func (s *SMap[KT, VT]) unsafeDelete(index int, k KT) {
	if s.Bucket[index] == nil {
		return
	}
	pre := s.Bucket[index]
	if pre.k == k {
		pre = pre.n
		return
	}
	current := pre.n
	for current != nil && current.k != k {
		pre = pre.n
		current = current.n
	}

	if current == nil {
		return
	}

	pre.n = current.n
}

func (s *SMap[KT, VT]) hashIndex(k KT) int {
	bs, _ := json.Marshal(k)
	h := int(cityhash.CityHash64(bs, uint32(len(bs)))) % len(s.Bucket)
	index := int(math.Abs(float64(h)))
	return index
}
