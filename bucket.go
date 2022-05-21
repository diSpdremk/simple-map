package simple_map

import (
	"encoding/json"
	"math"
	"reflect"

	"github.com/zhenjl/cityhash"
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

func (s *SMap[KT, VT]) hashIndex(k KT) int {
	switch reflect.TypeOf(k).Kind() {
	case reflect.Int:
		return k.(int)
	case reflect.Int8:
		return int(k.(int8))
	case reflect.Int32:
		return int(k.(int32))
	case reflect.Int64:
		return int(k.(int64))
	}
	bs, _ := json.Marshal(k)
	h := int(cityhash.CityHash64(bs, uint32(len(bs)))) % len(s.Bucket)
	index := int(math.Abs(float64(h)))
	return index
}
