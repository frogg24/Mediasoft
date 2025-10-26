package stack

import (
	"testing"
)

func TestNewStack(t *testing.T) {
	stack := NewStack[int](5)
	if stack.head != -1 {
		t.Errorf("Индекс головы стека должен быть -1, а не %d", stack.head)
	}
	if len(stack.s) != 5 {
		t.Errorf("Длина стека должна быть 5, а не %d", len(stack.s))
	}
}

func TestPush(t *testing.T) {
	stackString := NewStack[string](3)

	stackString.Push("first")
	if stackString.head != 0 {
		t.Errorf("Индекс головы стека должен быть 0, а не %d", stackString.head)
	}
	if stackString.s[0] != "first" {
		t.Errorf("Элемент должен быть равен 'first', а не %s", stackString.s[0])
	}

	stackString.Push("second")
	if stackString.head != 1 {
		t.Errorf("Индекс головы стека должен быть 1, а не %d", stackString.head)
	}
	if stackString.s[1] != "second" {
		t.Errorf("Элемент должен быть равен 'second', а не %s", stackString.s[1])
	}
}

func TestPop(t *testing.T) {
	stack := NewStack[int](3)
	stack.Push(10)
	stack.Push(20)

	v, ok := stack.Pop()
	if !ok {
		t.Error("Ожидаось, что вернется true, вернулось false")
	}
	if v != 20 {
		t.Errorf("Ожидалось, что вернется 20, вернулось %d", v)
	}
	if stack.head != 0 {
		t.Errorf("Индекс головы стека должен быть 0, а не %d", stack.head)
	}

	v, ok = stack.Pop()
	if !ok {
		t.Error("Ожидаось, что вернется true, вернулось false")
	}
	if v != 10 {
		t.Errorf("Ожидалось, что вернется 10, вернулось %d", v)
	}
	if stack.head != -1 {
		t.Errorf("Индекс головы стека должен быть -1, а не%d", stack.head)
	}

	// Попробуем Pop на пустом стеке
	_, ok = stack.Pop()
	if ok {
		t.Error("Ожидаось, что вернется false, вернулось true")
	}
}

func TestPeek(t *testing.T) {
	stack := NewStack[string](3)
	if stack.Peek() != nil {
		t.Error("Ожидалось nil")
	}

	stack.Push("peek-test")
	if stack.Peek() != "peek-test" {
		t.Errorf("Ожидалось значение 'peek-test', а не %v", stack.Peek())
	}

	if stack.head != 0 {
		t.Errorf("Ожидалось, что индекс головы стека останется 0, а не изменится на %d", stack.head)
	}
}

func TestStackOverflow(t *testing.T) {
	stack := NewStack[int](2)
	stack.Push(1)
	stack.Push(2)
	if stack.head != 1 {
		t.Errorf("Ожидалось голова индекса будет 1, а не %d", stack.head)
	}

	stack.Push(3)
	if stack.head != 1 {
		t.Errorf("Ожидалось голова индекса будет 1, а не %d", stack.head)
	}
	if stack.Peek() != 2 {
		t.Errorf("Ожидалось элемент будет равен 2, а не %d", stack.Peek())
	}
}
