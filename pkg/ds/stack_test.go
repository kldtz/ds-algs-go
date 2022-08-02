package ds

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPush(t *testing.T) {
	stack := NewStack[int]()
	stack.Push(3).Push(2).Push(1)

	for i := 1; i < 4; i += 1 {
		val, succ := stack.Pop()
		assert.True(t, succ)
		assert.Equal(t, i, val)
	}

	assert.Equal(t, uint(0), stack.Len())
}
