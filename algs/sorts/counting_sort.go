package sorts

import "github.com/kldtz/ds-algs-go/algs"

func CountingSort[T any](xs []T, key func(x T) int) []T {
	min := algs.Min(xs, key)
	max := algs.Max(xs, key)
	// Compute histogram
	count := make([]int, max-min+1)
	for _, x := range xs {
		count[key(x)-min] += 1
	}
	// Aggregate counts
	for i := 1; i < len(count); i += 1 {
		count[i] += count[i-1]
	}
	// Write elements to new slice in sorted order
	ys := make([]T, len(xs))
	for i := len(xs) - 1; i >= 0; i -= 1 {
		x := xs[i]
		k := key(x)
		count[k-min] -= 1
		ys[count[k-min]] = x
	}
	return ys
}
