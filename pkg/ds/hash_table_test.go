package ds

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetGet(t *testing.T) {
	table := NewHashTable[int]()

	table.Set("a", 1)
	table.Set("b", 2)
	table.Set("c", 3)

	assert.Equal(t, 3, table.Size())
	val, succ := table.Get("a")
	assert.True(t, succ)
	assert.Equal(t, 1, val)

	val, succ = table.Get("c")
	assert.True(t, succ)
	assert.Equal(t, 3, val)

	val, succ = table.Get("d")
	assert.False(t, succ)
}

func TestRehashing(t *testing.T) {
	table := NewHashTable[int]()
	n := 20
	for i := 0; i < n; i += 1 {
		key := strconv.Itoa(i)
		table.Set(key, i)
	}
	assert.Equal(t, n, table.Size())
	for i := 0; i < n; i += 1 {
		key := strconv.Itoa(i)
		val, succ := table.Get(key)
		assert.True(t, succ)
		assert.Equal(t, i, val)
	}
}
