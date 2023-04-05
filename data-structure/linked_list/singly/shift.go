package singly

func (s *Singly[T]) Shift() (value T, ok bool) {
	if s.Head == nil {
		return value, false
	}

	current := s.Head
	value, ok = current.Value, true

	s.Head = current.Next
	s.length--

	return
}
