package algs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandomInput(t *testing.T) {
	input := []int{9, 5, 2, 3, 8, 4, 1, 7, 6}
	BubbleSort(input)
	assert.True(t, IsSorted(input))
}
