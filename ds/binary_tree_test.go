package ds

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInorder(t *testing.T) {
	NIL := &BinaryNode[int]{}
	tree := &BinaryNode[int]{value: 1,
		left: &BinaryNode[int]{value: 2,
			left:  &BinaryNode[int]{value: 4, left: NIL, right: NIL},
			right: &BinaryNode[int]{value: 5, left: NIL, right: NIL},
		},
		right: &BinaryNode[int]{value: 3,
			left:  &BinaryNode[int]{value: 6, left: NIL, right: NIL},
			right: &BinaryNode[int]{value: 7, left: NIL, right: NIL},
		},
	}
	it := NewLevelOrder(tree, NIL)
	for i := 1; it.HasNext(); i += 1 {
		node := it.Next()
		binaryNode := node.(*BinaryNode[int])
		assert.Equal(t, i, binaryNode.value)
	}
}
