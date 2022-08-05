package algs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsSorted(t *testing.T) {
	input := []int{1, 3, 3, 4}
	assert.True(t, IsSorted(input))

	input = []int{1, 2, 3, 2}
	assert.False(t, IsSorted(input))
}

func TestMin(t *testing.T) {
	min := Min([]int{3, 1, 5, 2}, func(x int) int { return x })
	assert.Equal(t, 1, min)
}

func TestMax(t *testing.T) {
	max := Max([]int{3, 1, 5, 2}, func(x int) int { return x })
	assert.Equal(t, 5, max)
}
