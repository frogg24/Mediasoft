package binaryTree

type Tree[T types] struct {
	head *node[T]
}
type types interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~float32 | ~float64
}
type node[T types] struct {
	left, right *node[T]
	v           T
}

func NewBinaryTree[T types]() *Tree[T] {
	return &Tree[T]{}
}

// add - добавление значения в дерево
func (t *Tree[T]) Add(v T) {
	newNode := new(node[T])
	newNode.v = v
	newNode.left = nil
	newNode.right = nil
	if t.head == nil { //если список пуст
		t.head = newNode
		return
	}
	node := t.head
	for {
		if v >= node.v { // если значение текущего меньше значения добавляемого узла, то спуск в правый узел
			if node.right == nil {
				node.right = newNode
				return
			}
			node = node.right
		} else { // иначе в левый узел
			if node.left == nil {
				node.left = newNode
				return
			}
			node = node.left
		}
	}
}

// remove - удаление значения из дерева
func (t *Tree[T]) Remove(v T) {
	if t == nil || t.head == nil {
		return
	}

	var parent *node[T]
	current := t.head
	isLeftChild := false

	// поиск узла и его родителя
	for current != nil && current.v != v {
		parent = current
		if v < current.v {
			current = current.left
			isLeftChild = true
		} else {
			current = current.right
			isLeftChild = false
		}
	}

	// такого значения в дереве нет
	if current == nil {
		return
	}

	// если узел - лист
	if current.left == nil && current.right == nil {
		if parent == nil { // значит не вошел в цикл => удаляется корень
			t.head = nil
		} else if isLeftChild {
			parent.left = nil
		} else {
			parent.right = nil
		}
		return
	}

	// если узел имеет только одного потомка
	if current.left == nil || current.right == nil {
		var child *node[T]
		if current.left != nil {
			child = current.left
		} else {
			child = current.right
		}

		if parent == nil { // значит не вошел в цикл => удаляется корень
			t.head = child
		} else if isLeftChild {
			parent.left = child
		} else {
			parent.right = child
		}
		return
	}

	// если узел имеет 2 потомка
	// ищу самый правый узел слева
	rNode := current.left
	var parentrNode *node[T] = current

	for rNode.right != nil {
		parentrNode = rNode
		rNode = rNode.right
	}
	current.v = rNode.v

	if parentrNode == current {
		parentrNode.left = rNode.left // правее ничего быть не может
	} else {
		parentrNode.right = rNode.left
	}
}

func processNode[T types](current *node[T], result []T) []T {
	if current.left != nil {
		result = processNode(current.left, result)
	}
	result = append(result, current.v)
	if current.right != nil {
		result = processNode(current.right, result)
	}
	return result
}

// values - получение отсортированного слайса значений из дерева
func (t *Tree[T]) Values() []T {
	var result []T
	if t != nil && t.head != nil {
		result = processNode(t.head, result)
	}
	return result
}
