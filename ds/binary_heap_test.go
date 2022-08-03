package ds

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeap(t *testing.T) {
	heap := NewMaxHeap[int]()
	heap.insert(3).insert(2).insert(4).insert(1).insert(5)

	assert.Equal(t, 5, heap.Size())
	val, succ := heap.get_extreme()
	assert.True(t, succ)
	assert.Equal(t, 5, val)

	for i := 5; i > 0; i -= 1 {
		val, succ = heap.extract_extreme()
		assert.True(t, succ)
		assert.Equal(t, i, val)
	}

	_, succ = heap.get_extreme()
	assert.False(t, succ)
	_, succ = heap.extract_extreme()
	assert.False(t, succ)
}
