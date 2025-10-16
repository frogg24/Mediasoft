package linkedList

type SinglyLinkedList[T any] struct {
	first *item[T]
	last  *item[T]
	size  int
}

type item[T any] struct {
	v    T
	next *item[T]
}

func NewLinkedList[T any]() *SinglyLinkedList[T] {
	return &SinglyLinkedList[T]{}
}

// add - добавление значения в связный список
func (l *SinglyLinkedList[T]) Add(v T) {
	newItem := new(item[T])
	newItem.v = v
	//если лист пустой
	if l.first == nil {
		l.first = newItem
		l.last = newItem
	} else {
		l.last.next = newItem
		l.last = newItem
	}
	l.size++
}

// get - получение значения по индексу из связанного списка
func (l *SinglyLinkedList[T]) Get(idx int) T {
	//если индекс выходит за границы листа или лист nil
	if l == nil || idx >= l.size || idx < 0 {
		var v T
		return v
	}
	item := l.first
	for i := 0; i < idx; i++ {
		item = item.next
	}
	return item.v
}

// remove - удаление значения по индексу из списка
func (l *SinglyLinkedList[T]) Remove(idx int) {
	//если индекс выходит за границы листа или лист nil
	if l == nil || idx >= l.size || idx < 0 {
		return
	}
	item := l.first

	//буду думать, что item, на который более ничего не ссылается будет удален сборщиком мусора

	if l.size == 1 { // если в списек один элемент, он будет отчищен
		l.first = nil
		l.last = nil
	} else {
		if idx == 0 { // если удаляется первый элемент
			l.first = item.next
		} else {
			for i := 0; i < idx-1; i++ { //нахождения элемента idx-1
				item = item.next
			}
			if idx == l.size-1 { //если удаляется последний элемент
				l.last = item
				item.next = nil
			} else {
				item.next = item.next.next
			}
		}
	}
	l.size--
}

// values - получение слайса значений из списка
func (l *SinglyLinkedList[T]) Values() []T {
	if l == nil || l.size == 0 {
		return nil
	}
	values := make([]T, 0, l.size)
	item := l.first
	for i := 0; i < l.size; i++ {
		values = append(values, item.v)
		item = item.next
	}
	return values
}
