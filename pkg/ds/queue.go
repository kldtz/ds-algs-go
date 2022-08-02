package ds

type Queue[T any] struct {
	*List[T]
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{NewList[T]()}
}

func (q *Queue[T]) Enqueue(v T) *Queue[T] {
	q.Append(v)
	return q
}

func (q *Queue[T]) Dequeue() (T, bool) {
	return q.RemoveFirst()
}

func (q *Queue[T]) Peek() (T, bool) {
	if q.head == nil {
		var zero T
		return zero, false
	}
	return q.head.value, true
}
