package sorts

import "golang.org/x/exp/constraints"

func gappedInsertionSort[T constraints.Ordered](xs []T, gap int) {
	for i := gap; i < len(xs); i += 1 {
		// Loop invariant: xs[i-gap, i-2*gap, ...] is sorted
		cur := xs[i]
		// Move larger elements by one gap to the right and insert current element at correct gapped position
		j := i
		for ; j >= gap && cur < xs[j-gap]; j -= gap {
			xs[j] = xs[j-gap]
		}
		xs[j] = cur
	}
}

func InsertionSort[T constraints.Ordered](xs []T) {
	gappedInsertionSort(xs, 1)
}

func Shellsort[T constraints.Ordered](xs []T) {
	gaps := []int{701, 301, 132, 57, 23, 10, 4, 1}
	for _, gap := range gaps {
		gappedInsertionSort(xs, gap)
	}
}
