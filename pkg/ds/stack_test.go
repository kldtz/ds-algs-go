package ds

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPush(t *testing.T) {
	stack := NewStack[int]()
	stack.push(3)
	stack.push(2)
	stack.push(1)

	for i := 1; i < 4; i += 1 {
		val, succ := stack.pop()
		assert.True(t, succ)
		assert.Equal(t, i, val)
	}

	assert.Equal(t, uint(0), stack.Len())
}
