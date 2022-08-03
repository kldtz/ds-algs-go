package sorts

import (
	"testing"

	"github.com/kldtz/ds-algs-go/algs"

	"github.com/stretchr/testify/assert"
)

func TestQuicksort(t *testing.T) {
	input := algs.RandIntSlice(20)
	Quicksort(input)
	assert.True(t, algs.IsSorted(input))
}

func TestIterativeQuicksort(t *testing.T) {
	input := algs.RandIntSlice(20)
	IterativeQuicksort(input)
	assert.True(t, algs.IsSorted(input))
}

func TestTailRecursiveQuicksort(t *testing.T) {
	input := algs.RandIntSlice(20)
	TailRecursiveQuicksort(input)
	assert.True(t, algs.IsSorted(input))
}
