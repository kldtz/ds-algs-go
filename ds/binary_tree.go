package ds

type Node interface {
	Left() Node
	Right() Node
}

type BinaryNode[T any] struct {
	value T
	left  *BinaryNode[T]
	right *BinaryNode[T]
}

func (n *BinaryNode[T]) Left() Node {
	return n.left
}

func (n *BinaryNode[T]) Right() Node {
	return n.right
}

type TreeIterator interface {
	HasNext() bool
	Next() Node
}

type Inorder struct {
	cur   Node
	stack *Stack[Node]
	NIL   Node
}

func NewInorder(root Node, sentinel Node) *Inorder {
	return &Inorder{cur: root, stack: NewStack[Node](), NIL: sentinel}
}

func (it *Inorder) HasNext() bool {
	return !it.stack.IsEmpty() || it.cur != it.NIL
}

func (it *Inorder) Next() Node {
	if it.cur != it.NIL {
		cur := it.cur
		it.stack.Push(cur)
		it.cur = cur.Left()
		return cur
	} else {
		cur, _ := it.stack.Pop()
		it.cur = cur.Right()
		return cur
	}
}

type Preorder struct {
	cur   Node
	stack *Stack[Node]
	NIL   Node
}

func NewPreorder(root Node, sentinel Node) *Preorder {
	return &Preorder{cur: root, stack: NewStack[Node](), NIL: sentinel}
}

func (it *Preorder) HasNext() bool {
	return !it.stack.IsEmpty() || it.cur != it.NIL
}

func (it *Preorder) Next() Node {
	if it.cur != it.NIL {
		cur := it.cur
		it.stack.Push(cur)
		it.cur = cur.Left()
		return cur
	} else {
		cur, _ := it.stack.Pop()
		it.cur = cur.Right()
		return cur
	}
}

type Postorder struct {
	cur   Node
	prev  Node
	stack *Stack[Node]
	NIL   Node
}

func NewPostorder(root Node, sentinel Node) *Postorder {
	return &Postorder{cur: root, stack: NewStack[Node](), NIL: sentinel}
}

func (it *Postorder) HasNext() bool {
	return !it.stack.IsEmpty() || it.cur != it.NIL
}

func (it *Postorder) Next() Node {
	if it.cur != it.NIL {
		cur := it.cur
		it.stack.Push(cur)
		it.cur = cur.Left()
		return cur
	} else {
		cur, _ := it.stack.Peek()
		if cur.Right() == it.NIL || cur.Right() == it.prev {
			cur, _ = it.stack.Pop()
			it.prev = cur
			it.cur = it.NIL
			return cur
		} else {
			it.cur = it.cur.Right()
			return it.Next()
		}
	}
}

type LevelOrder struct {
	queue *Queue[Node]
	NIL   Node
}

func NewLevelOrder(root Node, sentinel Node) *LevelOrder {
	queue := NewQueue[Node]()
	queue.Enqueue(root)
	return &LevelOrder{queue: queue, NIL: sentinel}
}

func (it *LevelOrder) HasNext() bool {
	return !it.queue.IsEmpty()
}

func (it *LevelOrder) Next() Node {
	cur, _ := it.queue.Dequeue()
	if cur.Left() != it.NIL {
		it.queue.Enqueue(cur.Left())
	}
	if cur.Right() != it.NIL {
		it.queue.Enqueue(cur.Right())
	}
	return cur
}
