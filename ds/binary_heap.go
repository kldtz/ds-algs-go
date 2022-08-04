package ds

import (
	"github.com/kldtz/ds-algs-go/algs"
	"golang.org/x/exp/constraints"
)

type BinaryHeap[T constraints.Ordered] struct {
	xs   []T
	size int
	cmp  func(a T, b T) bool
}

func NewMaxHeap[T constraints.Ordered](xs []T) *BinaryHeap[T] {
	heap := &BinaryHeap[T]{
		xs:   xs,
		size: len(xs),
		cmp: func(a T, b T) bool {
			return a > b
		},
	}
	for i := heap.size >> 1; i >= 0; i -= 1 {
		heap.heapify(i)
	}
	return heap
}

func swap[T any](xs []T, i int, j int) {
	xs[i], xs[j] = xs[j], xs[i]
}

func parent(i int) int {
	return (i - 1) >> 1
}

func leftChild(i int) int {
	return i<<1 | 1
}

func rightChild(i int) int {
	return (i + 1) << 1
}

func (heap *BinaryHeap[T]) IsEmpty() bool {
	return heap.size == 0
}

func (heap *BinaryHeap[T]) Size() int {
	return heap.size
}

func (heap *BinaryHeap[T]) heapify(i int) {
	left := leftChild(i)
	right := rightChild(i)
	extreme := i
	if left < heap.size && heap.cmp(heap.xs[left], heap.xs[extreme]) {
		extreme = left
	}
	if right < heap.size && heap.cmp(heap.xs[right], heap.xs[extreme]) {
		extreme = right
	}
	if extreme != i {
		swap(heap.xs, i, extreme)
		heap.heapify(extreme)
	}
}

func (heap *BinaryHeap[T]) get_extreme() (T, bool) {
	if heap.size == 0 {
		var zero T
		return zero, false
	}
	return heap.xs[0], true
}

func (heap *BinaryHeap[T]) extract_extreme() (T, bool) {
	if heap.size == 0 {
		var zero T
		return zero, false
	}
	extreme := heap.xs[0]
	heap.xs[0] = heap.xs[heap.size-1]
	heap.size -= 1
	heap.heapify(0)
	return extreme, true
}

func (heap *BinaryHeap[T]) insert(v T) *BinaryHeap[T] {
	if heap.size < len(heap.xs) {
		heap.xs[heap.size] = v
	} else {
		heap.xs = append(heap.xs, v)
	}
	for i, p := heap.size, parent(heap.size); i != 0 && heap.xs[p] < heap.xs[i]; i, p = p, parent(p) {
		swap(heap.xs, i, p)
	}
	heap.size += 1
	return heap
}

func Heapsort[T constraints.Ordered](xs []T) {
	heap := NewMaxHeap(xs)
	for !heap.IsEmpty() {
		algs.Swap(heap.xs, 0, heap.size-1)
		heap.size -= 1
		heap.heapify(0)
	}
}
