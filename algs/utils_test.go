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
