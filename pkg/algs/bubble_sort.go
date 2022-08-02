package algs

import "golang.org/x/exp/constraints"

func BubbleSort[T constraints.Ordered](xs []T) {
	for n, last_swap := len(xs), 1; n > 1; n = last_swap {
		for i := 1; i < n; i += 1 {
			if xs[i-1] > xs[i] {
				Swap(xs, i-1, i)
				last_swap = i
			}
		}
	}
}
