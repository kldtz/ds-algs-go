package sorts

import "golang.org/x/exp/constraints"

func merge[T constraints.Ordered](left []T, right []T) []T {
	ys := make([]T, 0, len(left)+len(right))
	l, r := 0, 0
	for l < len(left) && r < len(right) {
		if left[l] <= right[r] {
			ys = append(ys, left[l])
			l += 1
		} else {
			ys = append(ys, right[r])
			r += 1
		}
	}
	for ; l < len(left); l += 1 {
		ys = append(ys, left[l])
	}
	for ; r < len(right); r += 1 {
		ys = append(ys, right[r])
	}
	return ys
}

func MergeSort[T constraints.Ordered](xs []T) []T {
	if len(xs) < 2 {
		return xs
	}
	mid := len(xs) / 2
	return merge(MergeSort(xs[:mid]), MergeSort(xs[mid:]))
}
