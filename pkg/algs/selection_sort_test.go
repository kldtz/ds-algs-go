package algs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSelectionSort(t *testing.T) {
	input := RandIntSlice(20)
	SelectionSort(input)
	assert.True(t, IsSorted(input))
}
