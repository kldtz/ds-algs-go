package sorts

import (
	"sort"
	"testing"

	"github.com/kldtz/ds-algs-go/algs"
	"github.com/stretchr/testify/assert"
)

func TestCountingSort(t *testing.T) {
	input := algs.RandIntSlice(20)
	actual := CountingSort(input, func(x int) int {
		return x
	})
	sort.Ints(input)
	assert.Equal(t, input, actual)
}
