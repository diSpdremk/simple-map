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

func (s *SMap[KT, VT]) Delete(k KT) {
	index := s.hashIndex(k)
	s.lock(index)
	s.unsafeDelete(index, k)
	s.unlock(index)
}

func (s *SMap[KT, VT]) Keys() []KT {
	var kts []KT
	for index, _ := range s.Bucket {
		if s.Bucket[index] == nil {
			continue // 并发不安全
		}
		s.lock(index)
		curr := s.Bucket[index]
		for curr != nil {
			kts = append(kts, curr.k)
			curr = curr.n
		}
		s.unlock(index)
	}
	return kts
}

func (s *SMap[KT, VT]) Values() []VT {
	var vts []VT
	for index, _ := range s.Bucket {
		if s.Bucket[index] == nil {
			continue // 并发不安全
		}
		s.lock(index)
		curr := s.Bucket[index]
		for curr != nil {
			vts = append(vts, curr.v)
			curr = curr.n
		}
		s.unlock(index)
	}
	return vts
}
