package sorts

import (
	"testing"

	"github.com/kldtz/ds-algs-go/algs"

	"github.com/stretchr/testify/assert"
)

func TestInsertionSort(t *testing.T) {
	input := algs.RandIntSlice(20)
	InsertionSort(input)
	assert.True(t, algs.IsSorted(input))
}

func TestShellsort(t *testing.T) {
	input := algs.RandIntSlice(20)
	Shellsort(input)
	assert.True(t, algs.IsSorted(input))
}
