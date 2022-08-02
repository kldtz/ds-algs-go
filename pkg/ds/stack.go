package ds

type Stack[T any] struct {
	*List[T]
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{NewList[T]()}
}

func (stack *Stack[T]) Push(v T) {
	stack.Prepend(v)
}

func (stack *Stack[T]) Peek() (T, bool) {
	if stack.head == nil {
		var zero T
		return zero, false
	}
	return stack.head.value, true
}

func (stack *Stack[T]) Pop() (T, bool) {
	return stack.RemoveFirst()
}
