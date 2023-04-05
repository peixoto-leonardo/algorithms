package singly

import "github.com/peixoto-leonardo/algorithms/data-structure/linked_list"

func (s *Singly[T]) GetLast() *linked_list.Node[T] {
	node := s.Head

	for node.Next != nil {
		node = node.Next
	}

	return node
}
