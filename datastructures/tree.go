package datastructures

import (
	"errors"
)

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

func (n *TreeNode) Insert(Value int) error {
	if n == nil {
		return errors.New("Cannot insert nil into tree")
	}

	switch {
	// Value already exists
	case Value == n.Value:
		return nil

	// lower val goes on the left child
	case Value < n.Value:
		// make leaf node
		if n.Left == nil {
			n.Left = &TreeNode{Value: Value}
			return nil
		}

		// if has child, recursively go down
		return n.Left.Insert(Value)

	// higher val goes on the right
	case Value > n.Value:
		// make leaf node
		if n.Right == nil {
			n.Right = &TreeNode{Value: Value}
			return nil
		}

		// if has child, recursively go down
		return n.Right.Insert(Value)
	}

	return nil
}

func (n *TreeNode) Find(s int) *TreeNode {
	if n == nil {
		return nil
	}

	switch {
	case s == n.Value:
		return n
	case s > n.Value:
		return n.Left.Find(s)
	default:
		return n.Right.Find(s)
	}
}

func (n *TreeNode) Delete(v int, parent *TreeNode) error {
	if n == nil {
		return errors.New("Value to be deleted does not exist in the tree")
	}

	switch {
	case v < n.Value:
		return n.Left.Delete(v, n)
	case v > n.Value:
		return n.Right.Delete(v, n)
	default:
		if n.Left == nil && n.Right == nil {
			n.replaceNode(parent, nil)
		} else if n.Left == nil {
			n.replaceNode(parent, n.Right)
		} else if n.Right == nil {
			n.replaceNode(parent, n.Left)
		} else {
			replacementNode, parentOfReplacementNode := n.findLargestChildNode(n)

			n.Value = replacementNode.Value

			return replacementNode.Delete(replacementNode.Value, parentOfReplacementNode)
		}
	}

	return nil
}

func (n *TreeNode) findLargestChildNode(parent *TreeNode) (*TreeNode, *TreeNode) {
	if n.Right == nil {
		return n, parent
	}
	return n.Right.findLargestChildNode(n)
}

func (n *TreeNode) replaceNode(parent, newNode *TreeNode) error {
	if n == nil {
		return errors.New("replaceNode() not allowed on a nil node")
	}

	if n == parent.Left {
		parent.Left = newNode
	} else {
		parent.Right = newNode
	}

	return nil
}

type Tree struct {
	Root *TreeNode
}

func (t *Tree) Insert(value int) error {
	if t.Root == nil {
		t.Root = &TreeNode{
			Value: value,
		}
		return nil
	}
	t.Root.Insert(value)
	return nil
}

func (t *Tree) InOrder() <-chan int {
	ch := make(chan int)
	go func() {
		t.inOrder(t.Root, ch)
		close(ch)
	}()
	return ch
}

func (t *Tree) inOrder(node *TreeNode, ch chan int) {
	if node != nil {
		t.inOrder(node.Left, ch)
		ch <- node.Value
		t.inOrder(node.Right, ch)
	}
}

func (t *Tree) PreOrder() <-chan int {
	ch := make(chan int)
	go func() {
		t.preOrder(t.Root, ch)
		close(ch)
	}()
	return ch
}

func (t *Tree) preOrder(node *TreeNode, ch chan int) {
	if node != nil {
		ch <- node.Value
		t.preOrder(node.Left, ch)
		t.preOrder(node.Right, ch)
	}
}

func (t *Tree) PostOrder() <-chan int {
	ch := make(chan int)
	go func() {
		t.postOrder(t.Root, ch)
		close(ch)
	}()
	return ch
}

func (t *Tree) postOrder(node *TreeNode, ch chan int) {
	if node != nil {
		t.postOrder(node.Left, ch)
		t.postOrder(node.Right, ch)
		ch <- node.Value
	}
}

func (t *Tree) Height() int {
	if t == nil {
		return 0
	}

	return t.Root.height()
}

func (n *TreeNode) height() int {
	if n == nil {
		return 0
	}

	return 1 + max(n.Left.height(), n.Right.height())
}

func (t *Tree) Diameter() int {
	if t == nil {
		return 0
	}

	leftHeight := t.Root.Left.height()
	rightHeight := t.Root.Right.height()

	leftTreeDiameter := t.Root.Left.diameter()
	rightTreeDiameter := t.Root.Right.diameter()

	return max(leftHeight+rightHeight+1, max(leftTreeDiameter, rightTreeDiameter))
}

func (n *TreeNode) diameter() int {
	if n == nil {
		return 0
	}

	leftHeight := n.Left.height()
	rightHeight := n.Right.height()

	leftTreeDiameter := n.Left.diameter()
	rightTreeDiameter := n.Right.diameter()

	return max(leftHeight+rightHeight+1, max(leftTreeDiameter, rightTreeDiameter))
}

func (n *Tree) LowestCommonAncestor(a, b int) *TreeNode {
	return n.Root.lowestCommonAncestorRecursive(a, b)
}

func (n *TreeNode) lowestCommonAncestorRecursive(a, b int) *TreeNode {
	if n == nil {
		return nil
	}

	if a > n.Value && b > n.Value {
		return n.Right.lowestCommonAncestorRecursive(a, b)
	}

	if a < n.Value && b < n.Value {
		return n.Left.lowestCommonAncestorRecursive(a, b)
	}

	return n
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
