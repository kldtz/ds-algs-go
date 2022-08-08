package sorting

import (
	"golang.org/x/exp/constraints"
)

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

func StackBasedInsertionSort[T constraints.Ordered](xs []T) []T {
	asc := make([]T, 0, len(xs))
	desc := make([]T, 0, len(xs))
	for _, x := range xs {
		if len(asc) == 0 || x > asc[len(asc)-1] {
			for len(desc) != 0 && x > desc[len(desc)-1] {
				asc = append(asc, desc[len(desc)-1])
				desc = desc[:len(desc)-1]
			}
			desc = append(desc, x)
		} else {
			for len(asc) != 0 && x < asc[len(asc)-1] {
				desc = append(desc, asc[len(asc)-1])
				asc = asc[:len(asc)-1]
			}
			asc = append(asc, x)
		}
	}
	for i := len(desc) - 1; i >= 0; i -= 1 {
		asc = append(asc, desc[i])
	}
	return asc
}
