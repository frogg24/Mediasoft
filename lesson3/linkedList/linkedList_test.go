package linkedList

import (
	"slices"
	"testing"
)

func TestNewLinkedList(t *testing.T) {
	l := NewLinkedList[int]()
	if l.first != nil {
		t.Error("Ожидается, что first == nil")
	}
	if l.last != nil {
		t.Error("Ожидается, что last == nil")
	}
	if l.size != 0 {
		t.Errorf("Ожидается, что size == 0, а не %d", l.size)
	}
}

func TestAdd(t *testing.T) {
	l := NewLinkedList[string]()
	l.Add("first")
	if l.size != 1 {
		t.Errorf("Ожидается, что size == 1 после добавления одного элемента, а не %d", l.size)
	}
	if l.first == nil {
		t.Error("Ожидается, что first != nil после добавления элемента")
	}
	if l.last == nil {
		t.Error("Ожидается, что last != nil после добавления элемента")
	}
	if l.first.v != "first" {
		t.Errorf("Ожидается, что первый элемент == 'first', а не %v", l.first.v)
	}
	if l.last.v != "first" {
		t.Errorf("Ожидается, что последний элемент == 'first', а не %v", l.last.v)
	}

	l.Add("second")
	if l.size != 2 {
		t.Errorf("Ожидается, что size == 2, а не %d", l.size)
	}
	if l.first.v != "first" {
		t.Errorf("Ожидается, что первый элемент == 'first', а не %v", l.first.v)
	}
	if l.last.v != "second" {
		t.Errorf("Ожидается, что последний элемент == 'second', а не %v", l.last.v)
	}
}

func TestGet(t *testing.T) {
	l := NewLinkedList[int]()
	l.Add(10)
	l.Add(20)
	l.Add(30)

	tests := []struct {
		name     string
		idx      int
		expected int
	}{
		{name: "Получить первый элемент", idx: 0, expected: 10},
		{name: "Получить второй элемент", idx: 1, expected: 20},
		{name: "Получить третий элемент", idx: 2, expected: 30},
		{name: "Индекс за границами (больше)", idx: 5, expected: 0},
		{name: "Индекс за границами (меньше)", idx: -1, expected: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := l.Get(tt.idx)
			if result != tt.expected {
				t.Errorf("Ожидается, что l.Get(%d) == %v, получено %v", tt.idx, tt.expected, result)
			}
		})
	}
}

func TestRemove(t *testing.T) {
	tests := []struct {
		name        string
		setup       func() *SinglyLinkedList[int]
		removeIdx   int
		expected    []int
		expectedLen int
	}{
		{
			name: "Удаление первого элемента",
			setup: func() *SinglyLinkedList[int] {
				l := NewLinkedList[int]()
				l.Add(10)
				l.Add(20)
				l.Add(30)
				return l
			},
			removeIdx:   0,
			expected:    []int{20, 30},
			expectedLen: 2,
		},
		{
			name: "Удаление последнего элемента",
			setup: func() *SinglyLinkedList[int] {
				l := NewLinkedList[int]()
				l.Add(10)
				l.Add(20)
				l.Add(30)
				return l
			},
			removeIdx:   2,
			expected:    []int{10, 20},
			expectedLen: 2,
		},
		{
			name: "Удаление среднего элемента",
			setup: func() *SinglyLinkedList[int] {
				l := NewLinkedList[int]()
				l.Add(10)
				l.Add(20)
				l.Add(30)
				return l
			},
			removeIdx:   1,
			expected:    []int{10, 30},
			expectedLen: 2,
		},
		{
			name: "Удаление единственного элемента",
			setup: func() *SinglyLinkedList[int] {
				l := NewLinkedList[int]()
				l.Add(10)
				return l
			},
			removeIdx:   0,
			expected:    nil,
			expectedLen: 0,
		},
		{
			name: "Удаление с индексом за границами (больше)",
			setup: func() *SinglyLinkedList[int] {
				l := NewLinkedList[int]()
				l.Add(10)
				l.Add(20)
				return l
			},
			removeIdx:   5,
			expected:    []int{10, 20},
			expectedLen: 2,
		},
		{
			name: "Удаление с индексом за границами (меньше)",
			setup: func() *SinglyLinkedList[int] {
				l := NewLinkedList[int]()
				l.Add(10)
				l.Add(20)
				return l
			},
			removeIdx:   -1,
			expected:    []int{10, 20},
			expectedLen: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := tt.setup()
			l.Remove(tt.removeIdx)

			if l.size != tt.expectedLen {
				t.Errorf("Ожидается, что размер списка == %d после удаления, а не %d", tt.expectedLen, l.size)
			}

			values := l.Values()
			if !slices.Equal(values, tt.expected) {
				t.Errorf("Ожидается, что значения == %v, а не %v", tt.expected, values)
			}
		})
	}
}

func TestValues(t *testing.T) {
	l := NewLinkedList[string]()
	if values := l.Values(); values != nil {
		t.Error("Ожидается, что Values() == nil для пустого списка")
	}

	l.Add("first")
	l.Add("second")
	l.Add("third")

	expected := []string{"first", "second", "third"}
	values := l.Values()
	if !slices.Equal(values, expected) {
		t.Errorf("Ожидается, что Values() == %v, а не %v", expected, values)
	}
}

func TestRemoveFromEmptyList(t *testing.T) {
	l := NewLinkedList[int]()
	l.Remove(0)
	if l.size != 0 {
		t.Errorf("Ожидается, что размер списка == 0 после попытки удаления из пустого списка, а не %d", l.size)
	}
}

func TestGetFromEmptyList(t *testing.T) {
	l := NewLinkedList[int]()
	result := l.Get(0)
	var zero int
	if result != zero {
		t.Errorf("Ожидается, что Get(0) возвращает нулевое значение для пустого списка, а не %v", result)
	}
}
