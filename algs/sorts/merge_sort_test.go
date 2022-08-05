package sorts

import (
	"sort"
	"testing"

	"github.com/kldtz/ds-algs-go/algs"
	"github.com/stretchr/testify/assert"
)

func TestMergeSort(t *testing.T) {
	input := algs.RandIntSlice(20)
	actual := MergeSort(input)
	sort.Ints(input)
	assert.Equal(t, input, actual)
}
