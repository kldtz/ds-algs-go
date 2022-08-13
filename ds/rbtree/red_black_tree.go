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

func (tree *RBTree[K, V]) Find(key K) (V, bool) {
	node, ok := tree.find(key)
	if ok {
		return node.Value, true
	}
	var zero V
	return zero, false
}

func (tree *RBTree[K, V]) find(key K) (*RBNode[K, V], bool) {
	node := tree.root
	for node != tree.NIL {
		if key < node.Key {
			node = node.left
		} else if key == node.Key {
			return node, true
		} else {
			node = node.right
		}
	}
	return nil, false
}

func (tree *RBTree[K, V]) Has(key K) bool {
	_, ok := tree.Find(key)
	return ok
}

func (tree *RBTree[K, V]) Delete(key K) bool {
	z, ok := tree.find(key)
	if !ok {
		return false
	}
	tree.deleteNode(z)
	tree.len -= 1
	return true
}

func (tree *RBTree[K, V]) deleteNode(z *RBNode[K, V]) {
	// maintain y as element to be moved or removed
	y := z
	yOriginalRed := y.red
	// x will be the start node for the fixup routine
	x := tree.NIL
	// if left child of z is empty
	if z.left == tree.NIL {
		// put right child in position of z
		x = z.right
		tree.transplant(z, z.right)
	} else if z.right == tree.NIL { // if right child is empty
		// put left child in position of z
		x = z.left
		tree.transplant(z, z.left)
	} else {
		// find smallest element greater than z
		y = tree.minimum(z.right)
		yOriginalRed = y.red
		x = y.right
		// if y is child of z
		if y.parent == z {
			x.parent = y
		} else {
			tree.transplant(y, y.right)
			y.right = z.right
			y.right.parent = y
		}
		// take left child from z and make it left child of y
		tree.transplant(z, y)
		y.left = z.left
		y.left.parent = y
		y.red = z.red
	}
	if !yOriginalRed {
		tree.fixupDelete(x)
	}
}

// Puts v in place of u
func (tree *RBTree[K, V]) transplant(u *RBNode[K, V], v *RBNode[K, V]) {
	if u.parent == tree.NIL { // u is the root node
		tree.root = v
	} else if u == u.parent.left { // u is a left child
		u.parent.left = v
	} else { // u is a right child
		u.parent.right = v
	}
	v.parent = u.parent
}

func (tree *RBTree[K, V]) minimum(node *RBNode[K, V]) *RBNode[K, V] {
	for node.left != tree.NIL {
		node = node.left
	}
	return node
}

func (tree *RBTree[K, V]) fixupDelete(x *RBNode[K, V]) {
	w := tree.NIL
	for x != tree.root && !x.red {
		// if x is a left child
		if x == x.parent.left {
			w = x.parent.right
			// case 1a: x's sibling is red
			if w.red {
				w.red = false
				x.parent.red = true
				tree.leftRotate(x.parent)
				w = x.parent.right
			}

			// case 2a: x's sibling w and both of w's children are black
			if !w.left.red && !w.right.red {
				w.red = true
				x = x.parent
			} else {
				// case 3a: x's sibling w is black, w's left child is red, w's right child is black
				if !w.right.red {
					w.left.red = false
					w.red = true
					tree.rightRotate(w)
					w = x.parent.right
				}
				// case 4a: x's sibling w is black, w's right child is red
				w.red = x.parent.red
				x.parent.red = false
				w.right.red = false
				tree.leftRotate(x.parent)
				x = tree.root
			}
		} else { // if x is a right child
			w = x.parent.left
			// case 1b: x's sibling is red
			if w.red {
				w.red = false
				x.parent.red = true
				tree.rightRotate(x.parent)
				w = x.parent.left
			}
			// case 2b: x's sibling w and both of w's children are black
			if !w.right.red && !w.left.red {
				w.red = true
				x = x.parent
			} else {
				if !w.left.red {
					w.right.red = false
					w.red = true
					tree.leftRotate(w)
					w = x.parent.left
				}
				w.red = x.parent.red
				x.parent.red = false
				w.left.red = false
				tree.rightRotate(x.parent)
				x = tree.root
			}
		}
	}
	x.red = false
}

type rbIterator[K constraints.Ordered, V any] interface {
	HasNext() bool
	Next() RBNode[K, V]
	Len() int
}

func Collect[K constraints.Ordered, V any, O any](it rbIterator[K, V], extractor func(RBNode[K, V]) O) []O {
	collection := make([]O, 0, it.Len())
	for it.HasNext() {
		collection = append(collection, extractor(it.Next()))
	}
	return collection
}

type LevelOrder[K constraints.Ordered, V any] struct {
	inner *ds.LevelOrder
	len   int
}

func (it *LevelOrder[K, V]) HasNext() bool {
	return it.inner.HasNext()
}

func (it *LevelOrder[K, V]) Next() RBNode[K, V] {
	node := it.inner.Next()
	rbNode := node.(*RBNode[K, V])
	return *rbNode
}

func (it *LevelOrder[K, V]) Len() int {
	return it.len
}

func (it *LevelOrder[K, V]) Keys() []K {
	keys := make([]K, 0, it.len)
	for it.HasNext() {
		keys = append(keys, it.Next().Key)
	}
	return keys
}

func (tree *RBTree[K, V]) Levelorder() *LevelOrder[K, V] {
	it := ds.NewLevelOrder(tree.root, tree.NIL)
	return &LevelOrder[K, V]{inner: it, len: tree.len}
}

type Inorder[K constraints.Ordered, V any] struct {
	inner *ds.Inorder
	len   int
}

func (it *Inorder[K, V]) HasNext() bool {
	return it.inner.HasNext()
}

func (it *Inorder[K, V]) Next() RBNode[K, V] {
	node := it.inner.Next()
	rbNode := node.(*RBNode[K, V])
	return *rbNode
}

func (it *Inorder[K, V]) Len() int {
	return it.len
}

func (tree *RBTree[K, V]) Inorder() *Inorder[K, V] {
	it := ds.NewInorder(tree.root, tree.NIL)
	return &Inorder[K, V]{inner: it, len: tree.len}
}
