package binaryTree

import (
	"slices"
	"testing"
)

func TestNewBinaryTree(t *testing.T) {
	tree := NewBinaryTree[int]()
	if tree.head != nil {
		t.Error("Ожидается, что head == nil")
	}
}

func TestAdd(t *testing.T) {
	tree := NewBinaryTree[int]()
	tree.Add(10)

	if tree.head == nil {
		t.Error("Ожидается, что head != nil")
	}
	if tree.head.v != 10 {
		t.Errorf("Ожидается, что значение корня == 10, а не %v", tree.head.v)
	}

	tree.Add(5)
	tree.Add(15)
	tree.Add(3)
	tree.Add(7)

	expected := []int{3, 5, 7, 10, 15}
	actual := tree.Values()
	if !slices.Equal(actual, expected) {
		t.Errorf("Ожидается, что Values() == %v, а не %v", expected, actual)
	}

	if tree.head.left.v != 5 {
		t.Errorf("Ожидается, что элемент равен 5, а не %v", tree.head.left.v)
	}
	if tree.head.right.v != 15 {
		t.Errorf("Ожидается, что элемент равен 5, а не %v", tree.head.right.v)
	}
	if tree.head.left.left.v != 3 {
		t.Errorf("Ожидается, что элемент равен 3, а не %v", tree.head.left.left.v)
	}
	if tree.head.left.right.v != 7 {
		t.Errorf("Ожидается, что элемент равен 7, а не %v", tree.head.left.right.v)
	}
}

func TestValues(t *testing.T) {
	tree := NewBinaryTree[float64]()
	if values := tree.Values(); len(values) != 0 {
		t.Error("Ожидается, что Values() возвращает пустой слайс для пустого дерева")
	}

	tree.Add(5.5)
	tree.Add(3.3)
	tree.Add(7.7)
	tree.Add(1.1)

	expected := []float64{1.1, 3.3, 5.5, 7.7}
	actual := tree.Values()
	if !slices.Equal(actual, expected) {
		t.Errorf("Ожидается, что Values() == %v, а не %v", expected, actual)
	}
}

func TestRemove(t *testing.T) {
	tests := []struct {
		name     string
		setup    func() *Tree[int]
		remove   int
		expected []int
	}{
		{
			name: "Удаление корня",
			setup: func() *Tree[int] {
				tree := NewBinaryTree[int]()
				tree.Add(10)
				return tree
			},
			remove:   10,
			expected: []int{},
		},
		{
			name: "Удаление листа слева",
			setup: func() *Tree[int] {
				tree := NewBinaryTree[int]()
				tree.Add(10)
				tree.Add(5)
				tree.Add(15)
				tree.Add(3)
				return tree
			},
			remove:   3,
			expected: []int{5, 10, 15},
		},
		{
			name: "Удаление листа справа",
			setup: func() *Tree[int] {
				tree := NewBinaryTree[int]()
				tree.Add(10)
				tree.Add(5)
				tree.Add(6)
				tree.Add(3)
				return tree
			},
			remove:   6,
			expected: []int{3, 5, 10},
		},
		{
			name: "Удаление узла с одним потомком слева",
			setup: func() *Tree[int] {
				tree := NewBinaryTree[int]()
				tree.Add(10)
				tree.Add(5)
				tree.Add(15)
				tree.Add(3)
				tree.Add(1)
				return tree
			},
			remove:   3,
			expected: []int{1, 5, 10, 15},
		},
		{
			name: "Удаление узла с одним потомком справа",
			setup: func() *Tree[int] {
				tree := NewBinaryTree[int]()
				tree.Add(10)
				tree.Add(5)
				tree.Add(15)
				tree.Add(3)
				tree.Add(4)
				return tree
			},
			remove:   3,
			expected: []int{1, 5, 10, 15},
		},
		{
			name: "Удаление узла с двумя потомками слева",
			setup: func() *Tree[int] {
				tree := NewBinaryTree[int]()
				tree.Add(10)
				tree.Add(5)
				tree.Add(15)
				tree.Add(3)
				tree.Add(7)
				tree.Add(12)
				tree.Add(18)
				return tree
			},
			remove:   10,
			expected: []int{3, 5, 7, 12, 15, 18},
		},
		{
			name: "Удаление узла с двумя потомками справа",
			setup: func() *Tree[int] {
				tree := NewBinaryTree[int]()
				tree.Add(10)
				tree.Add(5)
				tree.Add(15)
				tree.Add(3)
				return tree
			},
			remove:   10,
			expected: []int{3, 5, 15},
		},
		{
			name: "Удаление корня",
			setup: func() *Tree[int] {
				tree := NewBinaryTree[int]()
				tree.Add(10)
				tree.Add(5)
				tree.Add(15)
				tree.Add(3)
				tree.Add(7)
				tree.Add(12)
				tree.Add(18)
				return tree
			},
			remove:   10,
			expected: []int{3, 5, 7, 12, 15, 18},
		},
		{
			name: "Удаление корня с потомком справа",
			setup: func() *Tree[int] {
				tree := NewBinaryTree[int]()
				tree.Add(10)
				tree.Add(15)
				return tree
			},
			remove:   10,
			expected: []int{15},
		},
		{
			name: "Удаление узла с одним потомком справа",
			setup: func() *Tree[int] {
				tree := NewBinaryTree[int]()
				tree.Add(10)
				tree.Add(15)
				tree.Add(20)
				return tree
			},
			remove:   15,
			expected: []int{10, 20},
		},
		{
			name: "Удаление несуществующего узла",
			setup: func() *Tree[int] {
				tree := NewBinaryTree[int]()
				tree.Add(10)
				tree.Add(5)
				tree.Add(15)
				return tree
			},
			remove:   99,
			expected: []int{5, 10, 15},
		},
		{
			name: "Удаление единственного узла",
			setup: func() *Tree[int] {
				tree := NewBinaryTree[int]()
				tree.Add(42)
				return tree
			},
			remove:   42,
			expected: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := tt.setup()
			tree.Remove(tt.remove)
			actual := tree.Values()
			if !slices.Equal(actual, tt.expected) {
				t.Errorf("После удаления %d ожидается Values() == %v, получено %v", tt.remove, tt.expected, actual)
			}
		})
	}
}

func TestRemoveFromEmptyTree(t *testing.T) {
	tree := NewBinaryTree[int]()
	tree.Remove(10)
}
