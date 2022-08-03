package sorts

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
