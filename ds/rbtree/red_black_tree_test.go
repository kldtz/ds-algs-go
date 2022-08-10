package rbtree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateSingleElementTree(t *testing.T) {
	tree := NewRBTree[int, string]()
	tree.Insert(4, "four")
	assert.Equal(t, tree.len, 1)
	assert.Equal(t, tree.root.Key, 4)
}

func TestInsertionCase1aRight(t *testing.T) {
	tree := NewRBTree[int, string]()
	tree.Insert(3, "three").
		Insert(1, "one").
		Insert(4, "four").
		Insert(2, "two")

	it := tree.Levelorder()
	keys, colors := make([]int, 0, 4), make([]bool, 0, 4)
	for it.HasNext() {
		node := it.Next()
		keys = append(keys, node.Key)
		colors = append(colors, node.red)
	}
	assert.Equal(t, []int{3, 1, 4, 2}, keys)
	assert.Equal(t, []bool{false, false, false, true}, colors)
}

func TestInsertionCase1aLeft(t *testing.T) {
	tree := NewRBTree[int, string]()
	tree.Insert(3, "").
		Insert(2, "").
		Insert(4, "").
		Insert(1, "")

	it := tree.Levelorder()
	keys, colors := make([]int, 0, 4), make([]bool, 0, 4)
	for it.HasNext() {
		node := it.Next()
		keys = append(keys, node.Key)
		colors = append(colors, node.red)
	}
	assert.Equal(t, []int{3, 2, 4, 1}, keys)
	assert.Equal(t, []bool{false, false, false, true}, colors)
}
