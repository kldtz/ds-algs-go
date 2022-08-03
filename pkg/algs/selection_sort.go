package algs

import "golang.org/x/exp/constraints"

func SelectionSort[T constraints.Ordered](xs []T) {
	for i := 0; i < len(xs)-1; i += 1 {
		minIdx := i
		for j := i + 1; j < len(xs); j += 1 {
			if xs[j] < xs[minIdx] {
				minIdx = j
			}
		}
		Swap(xs, i, minIdx)
	}
}
