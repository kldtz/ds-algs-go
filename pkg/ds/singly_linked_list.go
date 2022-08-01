package ds

type Node[T any] struct {
	Value T
	Next  *Node[T]
}

type List[T any] struct {
	Head, Tail *Node[T]
	Len        uint
}

func NewList[T any]() *List[T] {
	return &List[T]{}
}

func (list *List[T]) ToSlice() []T {
	slice := make([]T, list.Len)
	for cur, i := list.Head, 0; cur != nil; cur, i = cur.Next, i+1 {
		slice[i] = cur.Value
	}
	return slice
}

func (list *List[T]) Append(v T) {
	node := &Node[T]{Value: v}
	if list.Head == nil {
		list.Head = node
		list.Tail = node
	} else {
		list.Tail.Next = node
		list.Tail = node
	}
	list.Len += 1
}

func (list *List[T]) Prepend(v T) {
	node := &Node[T]{Value: v}
	if list.Tail == nil {
		list.Head = node
		list.Tail = node
	} else {
		node.Next = list.Head
		list.Head = node
	}
	list.Len += 1
}

func FindFirst[T comparable](list *List[T], v T) (*Node[T], bool) {
	if list.Head == nil {
		return nil, false
	}
	for cur := list.Head; cur != nil; cur = cur.Next {
		if cur.Value == v {
			return cur, true
		}
	}
	return nil, false
}

func (list *List[T]) RemoveFirst() (*Node[T], bool) {
	if list.Head == nil {
		return nil, false
	}

	removed := list.Head
	list.Head = list.Head.Next
	if list.Head == nil {
		list.Tail = nil
	}
	list.Len -= 1
	return removed, true
}

func (list *List[T]) Delete(n *Node[T]) bool {
	if list.Head == nil {
		return false
	}

	if list.Head == n {
		list.Head = list.Head.Next
		list.Len -= 1
		return true
	}

	prev := list.Head
	for cur := prev.Next; cur != nil; prev, cur = cur, cur.Next {
		if cur == n {
			prev.Next = cur.Next
			list.Len -= 1
			return true
		}
	}
	return false
}

func (list *List[T]) Iterator() ListIterator[T] {
	return ListIterator[T]{cur: list.Head}
}

type ListIterator[T any] struct {
	cur *Node[T]
}

func (it *ListIterator[T]) HasNext() bool {
	return it.cur.Next != nil
}

func (it *ListIterator[T]) Next() T {
	curValue := it.cur.Value
	it.cur = it.cur.Next
	return curValue
}
