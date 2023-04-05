package singly

import "github.com/peixoto-leonardo/algorithms/data-structure/linked_list"

func (s *Singly[T]) Push(value T) {
	node := linked_list.NewNode(value)

	if s.Head == nil {
		s.Head = node
		s.length++
		return
	}

	s.GetLast().Next = node
	s.length++
}
