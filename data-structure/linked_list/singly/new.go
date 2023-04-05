package singly

import (
	"bytes"
	"fmt"

	"github.com/peixoto-leonardo/algorithms/data-structure/linked_list"
)

type Singly[T any] struct {
	length int
	Head   *linked_list.Node[T]
}

func New[T any](values ...T) *Singly[T] {
	singly := &Singly[T]{}

	for _, value := range values {
		singly.Push(value)
	}

	return singly
}

func (s *Singly[T]) String() string {
	var buffer bytes.Buffer

	for node := s.Head; node != nil; node = node.Next {
		buffer.WriteString(fmt.Sprint(node.Value))
	}

	return buffer.String()
}
