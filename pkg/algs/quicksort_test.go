package algs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuicksort(t *testing.T) {
	input := RandIntSlice(20)
	Quicksort(input)
	assert.True(t, IsSorted(input))
}

func TestIterativeQuicksort(t *testing.T) {
	input := RandIntSlice(20)
	IterativeQuicksort(input)
	assert.True(t, IsSorted(input))
}
