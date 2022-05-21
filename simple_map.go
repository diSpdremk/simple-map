package simple_map

type SMap[KT comparable, VT any] struct {
	Bucket    []*element[KT, VT]
	lockArray []uint32
}

func NewSimpleMap[KT comparable, VT any]() *SMap[KT, VT] {
	return &SMap[KT, VT]{
		Bucket:    make([]*element[KT, VT], SMapLen),
		lockArray: make([]uint32, SMapLen),
	}
}
