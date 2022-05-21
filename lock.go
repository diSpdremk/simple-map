package simple_map

import (
	"runtime"
	"sync/atomic"
)

// lock certain index of bucket
func (s *SMap[KT, VT]) lock(index int) {
	if index >= len(s.lockArray) {
		panic("illegal index")
	}
	for atomic.AddUint32(&s.lockArray[index], 1) != 1 {
		runtime.Gosched()
	}
}

// unlock certain index of bucket
func (s *SMap[KT, VT]) unlock(index int) {
	if index >= len(s.lockArray) {
		panic("illegal index")
	}
	atomic.StoreUint32(&s.lockArray[index], 0)
}
