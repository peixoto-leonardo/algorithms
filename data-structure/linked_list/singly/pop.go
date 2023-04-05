package singly

func (s *Singly[T]) Pop() (value T, ok bool) {
	if s.Head == nil {
		return value, ok
	}

	if s.Head.Next == nil {
		return s.Shift()
	}

	node := s.Head
	for node.Next.Next != nil {
		node = node.Next
		value, ok = node.Next.Value, true
	}
	node.Next = nil

	s.length--

	return
}
