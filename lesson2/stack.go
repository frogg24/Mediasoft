package main

type stack struct {
	s    []any // слайс в котором хранятся значения в стеке
	head int   // индекс головы стека
}

func newStack(size int) *stack {
	return &stack{
		s:    make([]any, size),
		head: -1,
	}
}

// push - добавление в стек значения
func push(s *stack, v any) {
	s.head++
	s.s[s.head] = v
}

// pop - получения значения из стека и его удаление из вершины
func pop(s *stack) any {
	if s.head == -1 {
		return nil
	}
	value := s.s[s.head]
	s.s[s.head] = nil
	s.head--
	return value
}

// peek - просмотр значения на вершине стека
func peek(s *stack) any {
	if s.head == -1 {
		return nil
	}
	return s.s[s.head]
}
