package core

type Stack[T any] struct {
	data []T
}

func New[T any]() *Stack[T] {
	return &Stack[T]{data: make([]T, 0)}
}

func (s *Stack[T]) Push(v T) {
	s.data = append(s.data, v)
}

func (s *Stack[T]) Pop() (T, bool) {
	var zero T
	if len(s.data) == 0 {
		return zero, false
	}
	i := len(s.data) - 1
	v := s.data[i]
	s.data = s.data[:i]
	return v, true
}

func (s *Stack[T]) Peek() (T, bool) {
	var zero T
	if len(s.data) == 0 {
		return zero, false
	}
	return s.data[len(s.data)-1], true
}

func (s *Stack[T]) Len() int {
	return len(s.data)
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.data) == 0
}
