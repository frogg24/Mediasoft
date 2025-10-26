package stack

type Stack[T any] struct {
	s    []T // слайс в котором хранятся значения в стеке
	head int // индекс головы стека
}

func NewStack[T any](size int) *Stack[T] {
	return &Stack[T]{
		s:    make([]T, size),
		head: -1,
	}
}

// push - добавление в стек значения
func (s *Stack[T]) Push(v T) {
	if s.head+2 > len(s.s) {
		return
	}
	s.head++
	s.s[s.head] = v
}

// pop - получения значения из стека и его удаление из вершины
func (s *Stack[T]) Pop() (T, bool) {
	if s.head < 0 {
		var v T
		return v, false
	}
	value := s.s[s.head]
	s.head--
	return value, true
}

// peek - просмотр значения на вершине стека
func (s *Stack[T]) Peek() any {
	if s.head == -1 {
		return nil
	}
	return s.s[s.head]
}
