package sorting

import (
	"github.com/kldtz/ds-algs-go/algs"
	"golang.org/x/exp/constraints"
)

func BubbleSort[T constraints.Ordered](xs []T) {
	for n := len(xs); n > 1; {
		last_swap := 1
		for i := 1; i < n; i += 1 {
			if xs[i-1] > xs[i] {
				algs.Swap(xs, i-1, i)
				last_swap = i
			}
		}
		n = last_swap
	}
}

func CocktailShakerSort[T constraints.Ordered](xs []T) {
	lo, hi := 0, len(xs)-1
	for lo < hi {
		new_hi := lo
		for i := lo; i < hi; i += 1 {
			if xs[i] > xs[i+1] {
				algs.Swap(xs, i, i+1)
				new_hi = i
			}
		}
		hi = new_hi

		new_lo := hi
		for i := hi; i > lo; i -= 1 {
			if xs[i-1] > xs[i] {
				algs.Swap(xs, i-1, i)
				new_lo = i
			}
		}
		lo = new_lo
	}
}

func CombSort[T constraints.Ordered](xs []T) {
	gap := len(xs)
	shrink := 1.3
	sorted := false

	for !sorted {
		gap = int(float64(gap) / shrink)
		if gap <= 1 {
			gap = 1
			sorted = true
		}

		for i := 0; i < len(xs)-gap; i += 1 {
			if xs[i] > xs[i+gap] {
				algs.Swap(xs, i, i+gap)
				sorted = false
			}
		}
	}
}
