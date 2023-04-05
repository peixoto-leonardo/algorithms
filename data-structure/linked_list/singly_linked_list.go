package linked_list

type Singly[T any] struct {
	length int
	Head   *Node[T]
}

func NewSingly[T any](values ...T) *Singly[T] {
	singly := &Singly[T]{}

	for _, value := range values {
		singly.Push(value)
	}

	return singly
}

func (s *Singly[T]) Unshift(value T) {
	node := NewNode(value)

	node.Next, s.Head = s.Head, node

	s.length++
}

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

func (s *Singly[T]) Push(value T) {
	node := NewNode(value)

	if s.Head == nil {
		s.Head = node
		s.length++
		return
	}

	s.GetLast().Next = node
	s.length++
}

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

func (s *Singly[T]) GetLast() *Node[T] {
	node := s.Head

	for node.Next != nil {
		node = node.Next
	}

	return node
}
