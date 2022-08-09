package ds

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func sentinelAndTree() (*BinaryNode[int], *BinaryNode[int]) {
	NIL := &BinaryNode[int]{}
	return NIL, &BinaryNode[int]{value: 1,
		left: &BinaryNode[int]{value: 2,
			left: &BinaryNode[int]{value: 4,
				left:  NIL,
				right: &BinaryNode[int]{value: 9, left: NIL, right: NIL}},
			right: &BinaryNode[int]{value: 5, left: NIL, right: NIL},
		},
		right: &BinaryNode[int]{value: 3,
			left: &BinaryNode[int]{value: 6, left: NIL, right: NIL},
			right: &BinaryNode[int]{value: 7,
				left:  &BinaryNode[int]{value: 14, left: NIL, right: NIL},
				right: NIL},
		},
	}
}

func TestLevelOrder(t *testing.T) {
	NIL, tree := sentinelAndTree()
	it := NewLevelOrder(tree, NIL)
	values := make([]int, 0, 9)
	for it.HasNext() {
		node := it.Next()
		binaryNode := node.(*BinaryNode[int])
		values = append(values, binaryNode.value)
	}
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 9, 14}, values)
}

func TestPreorder(t *testing.T) {
	NIL, tree := sentinelAndTree()
	it := NewPreorder(tree, NIL)
	values := make([]int, 0, 9)
	for it.HasNext() {
		node := it.Next()
		binaryNode := node.(*BinaryNode[int])
		values = append(values, binaryNode.value)
	}
	assert.Equal(t, []int{1, 2, 4, 9, 5, 3, 6, 7, 14}, values)
}

func TestInorder(t *testing.T) {
	NIL, tree := sentinelAndTree()
	it := NewInorder(tree, NIL)
	values := make([]int, 0, 9)
	for it.HasNext() {
		node := it.Next()
		binaryNode := node.(*BinaryNode[int])
		values = append(values, binaryNode.value)
	}
	assert.Equal(t, []int{4, 9, 2, 5, 1, 6, 3, 14, 7}, values)
}

func TestPostorder(t *testing.T) {
	NIL, tree := sentinelAndTree()
	it := NewPostorder(tree, NIL)
	values := make([]int, 0, 9)
	for it.HasNext() {
		node := it.Next()
		binaryNode := node.(*BinaryNode[int])
		values = append(values, binaryNode.value)
	}
	assert.Equal(t, []int{9, 4, 5, 2, 6, 14, 7, 3, 1}, values)
}
