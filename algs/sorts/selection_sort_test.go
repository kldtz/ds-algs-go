package sorts

import (
	"testing"

	"github.com/kldtz/ds-algs-go/algs"
	"github.com/stretchr/testify/assert"
)

func TestSelectionSort(t *testing.T) {
	input := algs.RandIntSlice(20)
	SelectionSort(input)
	assert.True(t, algs.IsSorted(input))
}
