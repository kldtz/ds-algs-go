package algs

import (
	"math/rand"

	"golang.org/x/exp/constraints"
)

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

func RandIntSlice(len int) []int {
	slice := make([]int, 0, len)
	for i := 0; i < len; i += 1 {
		slice = append(slice, rand.Intn(100))
	}
	return slice
}
