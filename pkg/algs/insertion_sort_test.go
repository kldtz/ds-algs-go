package algs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertionSort(t *testing.T) {
	input := RandIntSlice(20)
	InsertionSort(input)
	assert.True(t, IsSorted(input))
}

func TestShellsort(t *testing.T) {
	input := RandIntSlice(20)
	Shellsort(input)
	assert.True(t, IsSorted(input))
}
