package algs

import "golang.org/x/exp/constraints"

func Swap[T any](xs []T, i int, j int) {
	xs[i], xs[j] = xs[j], xs[i]
}

func IsSorted[T constraints.Ordered](xs []T) bool {
	for i := 1; i < len(xs); i += 1 {
		if xs[i] < xs[i-1] {
			return false
		}
	}
	return true
}
