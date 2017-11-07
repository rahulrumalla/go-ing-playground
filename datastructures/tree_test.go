package datastructures

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	sequence = []int{2, 5, 10, 14, 12, 7, 1, 11, 4, 0}
	tree     = &Tree{}
)

func init() {
	for _, v := range sequence {
		tree.Insert(v)
	}
}

func TestTreeInOrder(t *testing.T) {
	prev := 0
	for n := range tree.InOrder() {
		if n < 0 {
			break
		}
		fmt.Printf("Value: %d\n", n)
		assert.True(t, n > prev, "Wrong!")
		prev = n
	}
}

func TestTreePreOrder(t *testing.T) {
	for n := range tree.PreOrder() {
		if n < 0 {
			break
		}
		fmt.Printf("Value: %d\n", n)
	}
}

func TestTreePostOrder(t *testing.T) {
	for n := range tree.PostOrder() {
		if n < 0 {
			break
		}
		fmt.Printf("Value: %d\n", n)
	}
}

func TestTreeHeight(t *testing.T) {
	height := tree.Height()
	assert.Equal(t, 6, height, "Not same")
}

func TestTreeDiameter(t *testing.T) {
	diameter := tree.Diameter()
	assert.Equal(t, 8, diameter, "Not same")
}

func TestLowestCommonAncestor(t *testing.T) {
	var a, b int = 7, 12
	lcaNode := tree.LowestCommonAncestor(a, b)
	assert.Equal(t, 10, lcaNode.Value, "Not LCA")
}
