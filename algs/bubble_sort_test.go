package algs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSort(t *testing.T) {
	input := RandIntSlice(20)
	BubbleSort(input)
	assert.True(t, IsSorted(input))
}
