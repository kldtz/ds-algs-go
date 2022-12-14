package sorting

import (
	"testing"

	"github.com/kldtz/ds-algs-go/algs"
	"github.com/stretchr/testify/assert"
)

func TestBubbleSort(t *testing.T) {
	input := algs.RandIntSlice(20)
	BubbleSort(input)
	assert.True(t, algs.IsSorted(input))
}

func TestCocktailShakerSort(t *testing.T) {
	input := algs.RandIntSlice(20)
	CocktailShakerSort(input)
	assert.True(t, algs.IsSorted(input))
}

func TestCombSort(t *testing.T) {
	input := algs.RandIntSlice(20)
	CombSort(input)
	assert.True(t, algs.IsSorted(input))
}
