package rbtree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func getKey(node RBNode[int, string]) int {
	return node.Key
}

func getColor(node RBNode[int, string]) bool {
	return node.red
}

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

	keys := Collect[int, string](tree.Levelorder(), getKey)
	colors := Collect[int, string](tree.Levelorder(), getColor)
	assert.Equal(t, []int{3, 1, 4, 2}, keys)
	assert.Equal(t, []bool{false, false, false, true}, colors)
}

func TestInsertionCase1aLeft(t *testing.T) {
	tree := NewRBTree[int, string]()
	tree.Insert(3, "").
		Insert(2, "").
		Insert(4, "").
		Insert(1, "")

	keys := Collect[int, string](tree.Levelorder(), getKey)
	colors := Collect[int, string](tree.Levelorder(), getColor)
	assert.Equal(t, []int{3, 2, 4, 1}, keys)
	assert.Equal(t, []bool{false, false, false, true}, colors)
}

func TestInsertionCase1b(t *testing.T) {
	tree := NewRBTree[int, string]()
	tree.Insert(2, "").
		Insert(1, "").
		Insert(3, "").
		Insert(4, "")

	keys := Collect[int, string](tree.Levelorder(), getKey)
	colors := Collect[int, string](tree.Levelorder(), getColor)
	assert.Equal(t, []int{2, 1, 3, 4}, keys)
	assert.Equal(t, []bool{false, false, false, true}, colors)
}

func TestInsertionCase3a(t *testing.T) {
	tree := NewRBTree[int, string]()
	tree.Insert(3, "").
		Insert(2, "").
		Insert(1, "")

	keys := Collect[int, string](tree.Levelorder(), getKey)
	colors := Collect[int, string](tree.Levelorder(), getColor)
	assert.Equal(t, []int{2, 1, 3}, keys)
	assert.Equal(t, []bool{false, true, true}, colors)
}

func TestInsertionCase3b(t *testing.T) {
	tree := NewRBTree[int, string]()
	tree.Insert(1, "").
		Insert(2, "").
		Insert(3, "")

	keys := Collect[int, string](tree.Levelorder(), getKey)
	colors := Collect[int, string](tree.Levelorder(), getColor)
	assert.Equal(t, []int{2, 1, 3}, keys)
	assert.Equal(t, []bool{false, true, true}, colors)
}

func TestInsertionCase2And3(t *testing.T) {
	tree := NewRBTree[int, string]()
	tree.Insert(3, "").
		Insert(1, "").
		Insert(2, "")

	keys := Collect[int, string](tree.Levelorder(), getKey)
	colors := Collect[int, string](tree.Levelorder(), getColor)
	assert.Equal(t, []int{2, 1, 3}, keys)
	assert.Equal(t, []bool{false, true, true}, colors)
}

func TestInsertionCase2bAnd3b(t *testing.T) {
	tree := NewRBTree[int, string]()
	tree.Insert(1, "").
		Insert(3, "").
		Insert(2, "")

	keys := Collect[int, string](tree.Levelorder(), getKey)
	colors := Collect[int, string](tree.Levelorder(), getColor)
	assert.Equal(t, []int{2, 1, 3}, keys)
	assert.Equal(t, []bool{false, true, true}, colors)
}

func TestFind(t *testing.T) {
	tree := NewRBTree[int, string]()
	tree.Insert(1, "one").
		Insert(3, "three").
		Insert(2, "two").
		Insert(4, "four").
		Insert(5, "five")

	val, ok := tree.Find(1)
	assert.True(t, ok)
	assert.Equal(t, "one", val)
	val, ok = tree.Find(2)
	assert.True(t, ok)
	assert.Equal(t, "two", val)
	val, ok = tree.Find(3)
	assert.True(t, ok)
	assert.Equal(t, "three", val)
	val, ok = tree.Find(4)
	assert.True(t, ok)
	assert.Equal(t, "four", val)

	_, ok = tree.Find(6)
	assert.False(t, ok)
}

func TestHas(t *testing.T) {
	tree := NewRBTree[int, int]()
	tree.Insert(5, 5).Insert(2, 2).Insert(1, 1).Insert(3, 3)

	ok := tree.Has(5)
	assert.True(t, ok)
}

func TestSimpleDeletion(t *testing.T) {
	tree := NewRBTree[int, string]()
	tree.Insert(2, "").Insert(1, "").Insert(3, "")

	assert.Equal(t, tree.len, 3)

	ok := tree.Delete(3)
	assert.True(t, ok)
	assert.Equal(t, tree.len, 2)
	keys := Collect[int, string](tree.Levelorder(), getKey)
	assert.Equal(t, []int{2, 1}, keys)
}

func TestDeletionUnknownElement(t *testing.T) {
	tree := NewRBTree[int, string]()
	tree.Insert(2, "").Insert(1, "").Insert(3, "")
	assert.Equal(t, 3, tree.len)

	ok := tree.Delete(4)
	assert.False(t, ok)
	assert.Equal(t, 3, tree.len)
	keys := Collect[int, string](tree.Levelorder(), getKey)
	assert.Equal(t, []int{2, 1, 3}, keys)
}

func TestDeletionCase1a2a(t *testing.T) {
	tree := NewRBTree[int, string]()
	for i := 1; i < 7; i += 1 {
		tree.Insert(i, "")
	}
	assert.Equal(t, 6, tree.len)

	tree.Delete(1)
	assert.Equal(t, 5, tree.len)
	keys := Collect[int, string](tree.Levelorder(), getKey)
	assert.Equal(t, []int{4, 2, 5, 3, 6}, keys)
}

func TestDeletionCase3a4a(t *testing.T) {
	tree := NewRBTree[int, string]()
	tree.Insert(4, "").
		Insert(1, "").
		Insert(2, "").
		Insert(3, "")

	keys := Collect[int, string](tree.Levelorder(), getKey)
	assert.Equal(t, []int{2, 1, 4, 3}, keys)

	ok := tree.Delete(1)
	keys = Collect[int, string](tree.Levelorder(), getKey)
	assert.True(t, ok)
	assert.Equal(t, []int{3, 2, 4}, keys)
}

func TestDeleteTwoKeys3a4a2b(t *testing.T) {
	tree := NewRBTree[int, string]()
	tree.Insert(4, "").
		Insert(1, "").
		Insert(2, "").
		Insert(3, "")

	tree.Delete(1)
	tree.Delete(3)

	keys := Collect[int, string](tree.Levelorder(), getKey)
	assert.Equal(t, []int{4, 2}, keys)
}

func TestInorder(t *testing.T) {
	tree := NewRBTree[int, string]()
	tree.Insert(5, "5").Insert(2, "2").Insert(3, "3").Insert(1, "1").Insert(4, "4")

	keys := Collect[int, string](tree.Inorder(), getKey)
	assert.Equal(t, []int{1, 2, 3, 4, 5}, keys)

	tree.Delete(5)
	keys = Collect[int, string](tree.Inorder(), getKey)
	assert.Equal(t, []int{1, 2, 3, 4}, keys)

	tree.Delete(3)
	keys = Collect[int, string](tree.Inorder(), getKey)
	assert.Equal(t, []int{1, 2, 4}, keys)
}
