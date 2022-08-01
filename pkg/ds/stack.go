package ds

type Stack[T any] struct {
	list *List[T]
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{list: NewList[T]()}
}

func (stack *Stack[T]) push(v T) {
	stack.list.Prepend(v)
}

func (stack *Stack[T]) peek() (T, bool) {
	if stack.list.head == nil {
		var zero T
		return zero, false
	}
	return stack.list.head.value, true
}

func (stack *Stack[T]) pop() (T, bool) {
	return stack.list.RemoveFirst()
}

func (list *Stack[T]) Iterator() ListIterator[T] {
	return list.Iterator()
}

func (stack *Stack[T]) Len() uint {
	return stack.list.Len()
}
