package ds

import "golang.org/x/exp/constraints"

type node[K constraints.Ordered, V any] struct {
	key    K
	value  V
	red    bool
	parent *node[K, V]
	left   *node[K, V]
	right  *node[K, V]
}

func (n *node[K, V]) Left() Node {
	return n.left
}

func (n *node[K, V]) Right() Node {
	return n.right
}

type RedBlackTree[K constraints.Ordered, V any] struct {
	root *node[K, V]
	len  int
	NIL  *node[K, V] // sentinel: &node[K, V]{red: false}
}

func (tree *RedBlackTree[K, V]) Insert(key K, value V) {
	y := tree.NIL
	x := tree.root
	for x != tree.NIL {
		y = x
		if key < y.key {
			x = x.left
		} else if key == y.key {
			x.value = value
		} else {
			x = x.right
		}
	}
	z := &node[K, V]{key: key, value: value, red: true, parent: y, left: tree.NIL, right: tree.NIL}
	if y == tree.NIL {
		tree.root = z
	} else if key < y.key {
		y.left = z
	} else {
		y.right = z
	}
	tree.fixupInsert(z)
}

func (tree *RedBlackTree[K, V]) fixupInsert(z *node[K, V]) {
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
				z.red = false
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
func (tree *RedBlackTree[K, V]) leftRotate(x *node[K, V]) {
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
func (tree *RedBlackTree[K, V]) rightRotate(y *node[K, V]) {
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
