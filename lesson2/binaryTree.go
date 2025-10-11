package main

type tree struct {
	head *node
}

type node struct {
	left, right *node
	v           int
}

func newTree() *tree {
	return &tree{}
}

// add - добавление значения в дерево
func addToTree(t *tree, v int) {
	newNode := new(node)
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
func removeFromTree(t *tree, v int) {
	if t == nil || t.head == nil {
		return
	}

	var parent *node
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
		var child *node
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
	var parentrNode *node = current

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

func processNode(current *node, result []int) []int {
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
func values(t *tree) []int {
	var result []int
	result = processNode(t.head, result)
	return result
}
