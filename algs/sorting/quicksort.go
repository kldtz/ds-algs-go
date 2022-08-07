package sorting

import (
	"github.com/kldtz/ds-algs-go/algs"
	"github.com/kldtz/ds-algs-go/ds"

	"golang.org/x/exp/constraints"
)

func lomutoPartition[T constraints.Ordered](xs []T, lo int, hi int) int {
	pivot := xs[hi]
	boundary := lo
	// Loop invariant: all elements xs[:boundary] are smaller than the pivot
	for i := lo; i < hi; i += 1 {
		if xs[i] < pivot {
			algs.Swap(xs, i, boundary)
			boundary += 1
		}
	}
	algs.Swap(xs, boundary, hi)
	return boundary
}

func quicksort[T constraints.Ordered](xs []T, lo int, hi int) {
	if lo < hi {
		p := lomutoPartition(xs, lo, hi)
		quicksort(xs, lo, p-1)
		quicksort(xs, p+1, hi)
	}
}

func Quicksort[T constraints.Ordered](xs []T) {
	quicksort(xs, 0, len(xs)-1)
}

type loHi struct {
	lo int
	hi int
}

func IterativeQuicksort[T constraints.Ordered](xs []T) {
	stack := ds.NewStack[loHi]()
	stack.Push(loHi{0, len(xs) - 1})
	for !stack.IsEmpty() {
		cur, _ := stack.Pop()
		if cur.lo >= cur.hi {
			continue
		}
		p := lomutoPartition(xs, cur.lo, cur.hi)
		stack.Push(loHi{p + 1, cur.hi})
		stack.Push(loHi{cur.lo, p - 1})
	}
}

func tailRecursiveQuicksort[T constraints.Ordered](xs []T, lo int, hi int) {
	for lo < hi {
		// After partitioning, the pivot p is in its final position
		p := lomutoPartition(xs, lo, hi)
		tailRecursiveQuicksort(xs, lo, p-1)
		lo = p + 1
	}
}

func TailRecursiveQuicksort[T constraints.Ordered](xs []T) {
	tailRecursiveQuicksort(xs, 0, len(xs)-1)
}
