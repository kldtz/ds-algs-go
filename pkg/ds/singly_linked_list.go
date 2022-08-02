package ds

type Node[T any] struct {
	value T
	next  *Node[T]
}

type List[T any] struct {
	head, tail *Node[T]
	len        uint
}

func NewList[T any]() *List[T] {
	return &List[T]{}
}

func (list *List[T]) Len() uint {
	return list.len
}

func (list *List[T]) ToSlice() []T {
	slice := make([]T, list.len)
	for cur, i := list.head, 0; cur != nil; cur, i = cur.next, i+1 {
		slice[i] = cur.value
	}
	return slice
}

func (list *List[T]) Append(v T) *List[T] {
	node := &Node[T]{value: v}
	if list.head == nil {
		list.head = node
		list.tail = node
	} else {
		list.tail.next = node
		list.tail = node
	}
	list.len += 1
	return list
}

func (list *List[T]) Prepend(v T) *List[T] {
	node := &Node[T]{value: v}
	if list.tail == nil {
		list.head = node
		list.tail = node
	} else {
		node.next = list.head
		list.head = node
	}
	list.len += 1
	return list
}

func findFirst[T comparable](list *List[T], v T) (*Node[T], bool) {
	if list.head == nil {
		return nil, false
	}
	for cur := list.head; cur != nil; cur = cur.next {
		if cur.value == v {
			return cur, true
		}
	}
	return nil, false
}

func Contains[T comparable](list *List[T], v T) bool {
	_, found := findFirst(list, v)
	return found
}

func (list *List[T]) RemoveFirst() (T, bool) {
	if list.head == nil {
		var zero T
		return zero, false
	}

	removed := list.head.value
	list.head = list.head.next
	if list.head == nil {
		list.tail = nil
	}
	list.len -= 1
	return removed, true
}

func (list *List[T]) Delete(n *Node[T]) bool {
	if list.head == nil {
		return false
	}

	if list.head == n {
		list.head = list.head.next
		list.len -= 1
		return true
	}

	prev := list.head
	for cur := prev.next; cur != nil; prev, cur = cur, cur.next {
		if cur == n {
			prev.next = cur.next
			list.len -= 1
			return true
		}
	}
	return false
}

func (list *List[T]) Iterator() ListIterator[T] {
	return ListIterator[T]{cur: list.head}
}

type ListIterator[T any] struct {
	cur *Node[T]
}

func (it *ListIterator[T]) HasNext() bool {
	return it.cur.next != nil
}

func (it *ListIterator[T]) Next() T {
	curvalue := it.cur.value
	it.cur = it.cur.next
	return curvalue
}
