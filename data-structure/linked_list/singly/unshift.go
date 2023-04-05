package singly

import "github.com/peixoto-leonardo/algorithms/data-structure/linked_list"

func (s *Singly[T]) Unshift(value T) {
	node := linked_list.NewNode(value)

	node.Next, s.Head = s.Head, node

	s.length++
}
