package ds

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnqueue(t *testing.T) {
	queue := NewQueue[int]()
	queue.Enqueue(1).Enqueue(2).Enqueue(3)

	for i := 1; i < 4; i += 1 {
		val, succ := queue.Dequeue()
		assert.True(t, succ)
		assert.Equal(t, i, val)
	}

	assert.Equal(t, uint(0), queue.Len())
}
