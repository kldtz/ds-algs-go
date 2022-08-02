package ds

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAppend(t *testing.T) {
	list := NewList[int]()
	list.Append(1).Append(2)

	assert.Equal(t, uint(2), list.Len())
	assert.Equal(t, []int{1, 2}, list.ToSlice())
}

func TestPrepend(t *testing.T) {
	list := NewList[int]()
	list.Prepend(2).Prepend(1)

	assert.Equal(t, uint(2), list.Len())
	assert.Equal(t, []int{1, 2}, list.ToSlice())
}

func TestFindFirst(t *testing.T) {
	list := NewList[int]()
	list.Append(3).Append(2).Append(1)

	node, succ := findFirst(list, 2)
	assert.True(t, succ)
	assert.Equal(t, 2, node.value)
	assert.Equal(t, 1, node.next.value)
}

func TestFindFirstUnknown(t *testing.T) {
	list := NewList[int]()
	list.Append(3).Append(2).Append(1)

	node, succ := findFirst(list, 4)
	assert.Nil(t, node)
	assert.False(t, succ)
}

func TestRemoveFirst(t *testing.T) {
	list := NewList[int]()
	list.Append(3).Append(1).Append(2)

	val, succ := list.RemoveFirst()
	assert.Equal(t, true, succ)
	assert.Equal(t, 3, val)
	assert.Equal(t, uint(2), list.Len())
	assert.Equal(t, []int{1, 2}, list.ToSlice())
}

func TestDelete(t *testing.T) {
	list := NewList[int]()
	list.Append(3).Append(2).Append(1).Append(2)

	assert.Equal(t, uint(4), list.Len())
	node, _ := findFirst(list, 2)
	succ := list.Delete(node)
	assert.True(t, succ)
	assert.Equal(t, uint(3), list.Len())
	assert.Equal(t, []int{3, 1, 2}, list.ToSlice())
}

func TestIterator(t *testing.T) {
	list := NewList[int]()
	list.Append(1).Append(2).Append(3)

	it := list.Iterator()
	for i := 1; it.HasNext(); i += 1 {
		assert.Equal(t, i, it.Next())
	}
}
