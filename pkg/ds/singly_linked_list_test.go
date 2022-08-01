package ds

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAppend(t *testing.T) {
	list := NewList[int]()
	list.Append(1)
	list.Append(2)

	assert.Equal(t, uint(2), list.Len)
	assert.Equal(t, []int{1, 2}, list.ToSlice())
}

func TestPrepend(t *testing.T) {
	list := NewList[int]()
	list.Prepend(2)
	list.Prepend(1)

	assert.Equal(t, uint(2), list.Len)
	assert.Equal(t, []int{1, 2}, list.ToSlice())
}

func TestFindFirst(t *testing.T) {
	list := NewList[int]()
	list.Append(3)
	list.Append(2)
	list.Append(1)

	node, err := list.FindFirst(2)
	assert.Nil(t, err)
	assert.Equal(t, 2, node.Value)
	assert.Equal(t, 1, node.Next.Value)
}

func TestFindFirstUnknown(t *testing.T) {
	list := NewList[int]()
	list.Append(3)
	list.Append(2)
	list.Append(1)

	node, err := list.FindFirst(4)
	assert.Nil(t, node)
	assert.NotNil(t, err)
}

func TestRemoveFirst(t *testing.T) {
	list := NewList[int]()
	list.Append(3)
	list.Append(1)
	list.Append(2)

	list.RemoveFirst()
	assert.Equal(t, uint(2), list.Len)
	assert.Equal(t, []int{1, 2}, list.ToSlice())
}

func TestRemoveFirstMatching(t *testing.T) {
	list := NewList[int]()
	list.Append(3)
	list.Append(2)
	list.Append(1)
	list.Append(2)

	assert.Equal(t, uint(4), list.Len)
	node, err := list.RemoveFirstMatching(2)
	assert.Nil(t, err)
	assert.Equal(t, uint(3), list.Len)
	assert.Equal(t, 2, node.Value)
	assert.Equal(t, []int{3, 1, 2}, list.ToSlice())
}

func TestIterator(t *testing.T) {
	list := NewList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)

	it := list.Iterator()
	for i := 1; it.HasNext(); i += 1 {
		assert.Equal(t, i, it.Next())
	}
}
