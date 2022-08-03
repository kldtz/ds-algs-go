package sorts

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
