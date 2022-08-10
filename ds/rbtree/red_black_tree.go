package rbtree

import (
	"github.com/kldtz/ds-algs-go/ds"
	"golang.org/x/exp/constraints"
)

type RBNode[K constraints.Ordered, V any] struct {
	Key    K
	Value  V
	red    bool
	parent *RBNode[K, V]
	left   *RBNode[K, V]
	right  *RBNode[K, V]
}

func (n *RBNode[K, V]) Left() ds.Node {
	return n.left
}

func (n *RBNode[K, V]) Right() ds.Node {
	return n.right
}

type RBTree[K constraints.Ordered, V any] struct {
	root *RBNode[K, V]
	len  int
	NIL  *RBNode[K, V] // sentinel: &RBNode[K, V]{red: false}
}

func NewRBTree[K constraints.Ordered, V any]() *RBTree[K, V] {
	NIL := &RBNode[K, V]{red: false}
	return &RBTree[K, V]{
		root: NIL,
		len:  0,
		NIL:  NIL,
	}
}

func (tree *RBTree[K, V]) Insert(key K, value V) *RBTree[K, V] {
	y := tree.NIL
	x := tree.root
	for x != tree.NIL {
		y = x
		if key < y.Key {
			x = x.left
		} else if key == y.Key {
			x.Value = value
			return tree
		} else {
			x = x.right
		}
	}
	z := &RBNode[K, V]{Key: key, Value: value, red: true, parent: y, left: tree.NIL, right: tree.NIL}
	if y == tree.NIL {
		tree.root = z
	} else if key < y.Key {
		y.left = z
	} else {
		y.right = z
	}
	tree.fixupInsert(z)
	tree.len += 1
	return tree
}

func (tree *RBTree[K, V]) fixupInsert(z *RBNode[K, V]) {
	// while z's parent is red
	for z.parent.red {
		// if the parent of z is a left child
		if z.parent == z.parent.parent.left {
			// set y as the parent's right sibling
			y := z.parent.parent.right
			// case 1a: both the parent and its sibling are red
			if y.red {
				// color parent and sibling black
				z.parent.red = false
				y.red = false
				// color grandparent red
				z.parent.parent.red = true
				// continue with grandparent as z
				z = z.parent.parent
			} else { // sibling of the parent is black
				// case 2a: z is a right child
				if z == z.parent.right {
					z = z.parent
					tree.leftRotate(z)
				}
				// case 3a: z and its parent are left children
				z.parent.red = false
				z.parent.parent.red = true
				tree.rightRotate(z.parent.parent)
			}
		} else { // y is the parent's left sibling
			y := z.parent.parent.left
			// case 1b: both the parent and its sibling are red
			if y.red {
				// color parent and sibling black
				z.parent.red = false
				y.red = false
				// color grandparent red
				z.parent.parent.red = true
				// continue with grandparent as z
				z = z.parent.parent
			} else { // case 2b; z is left child
				if z == z.parent.left {
					z = z.parent
					tree.rightRotate(z)
				}
				// case 3b: z and its parent are right children
				z.parent.red = false
				z.parent.parent.red = true
				tree.leftRotate(z.parent.parent)
			}
		}
	}
	tree.root.red = false
}

//      x                y
//     / \              / \
//    a   y     =>     x   c
//       / \          / \
//      b   c        a   b
func (tree *RBTree[K, V]) leftRotate(x *RBNode[K, V]) {
	y := x.right
	x.right = y.left
	if y.left != tree.NIL {
		y.left.parent = x
	}
	y.parent = x.parent
	if x.parent == tree.NIL {
		tree.root = y
	} else if x == x.parent.left {
		x.parent.left = y
	} else {
		x.parent.right = y
	}
	y.left = x
	x.parent = y
}

//      x                y
//     / \              / \
//    a   y     <=     x   c
//       / \          / \
//      b   c        a   b
func (tree *RBTree[K, V]) rightRotate(y *RBNode[K, V]) {
	x := y.left
	y.left = x.right
	if x.right != tree.NIL {
		x.right.parent = y
	}
	x.parent = y.parent
	if y.parent == tree.NIL {
		tree.root = x
	} else if y == y.parent.right {
		y.parent.right = x
	} else {
		y.parent.left = x
	}
	x.right = y
	y.parent = x
}

type LevelOrder[K constraints.Ordered, V any] struct {
	inner *ds.LevelOrder
}

func (it *LevelOrder[K, V]) HasNext() bool {
	return it.inner.HasNext()
}

func (it *LevelOrder[K, V]) Next() RBNode[K, V] {
	node := it.inner.Next()
	rbNode := node.(*RBNode[K, V])
	return *rbNode
}

func (tree *RBTree[K, V]) Levelorder() *LevelOrder[K, V] {
	it := ds.NewLevelOrder(tree.root, tree.NIL)
	return &LevelOrder[K, V]{inner: it}
}

type Inorder[K constraints.Ordered, V any] struct {
	inner *ds.Inorder
}

func (it *Inorder[K, V]) HasNext() bool {
	return it.inner.HasNext()
}

func (it *Inorder[K, V]) Next() RBNode[K, V] {
	node := it.inner.Next()
	rbNode := node.(*RBNode[K, V])
	return *rbNode
}

func (tree *RBTree[K, V]) Inorder() *Inorder[K, V] {
	it := ds.NewInorder(tree.root, tree.NIL)
	return &Inorder[K, V]{inner: it}
}
