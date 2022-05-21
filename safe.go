package simple_map

func (s *SMap[KT, VT]) Get(k KT) (v VT, exist bool) {
	index := s.hashIndex(k)
	s.lock(index)
	v, exist = s.unsafeGet(index, k)
	s.unlock(index)
	return
}

func (s *SMap[KT, VT]) Set(k KT, v VT) {
	index := s.hashIndex(k)
	s.lock(index)
	s.unsafeSet(index, k, v)
	s.unlock(index)
}
