package datastructures

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStackArray(t *testing.T) {
	stack := NewStackArray()
	stack.Push("One")
	stack.Push("Two")
	stack.Push("Three")

	assert.Equal(t, 3, stack.Size(), "Not equal")

	item := stack.Pop()
	assert.Equal(t, "Three", item, "Not equal")

	stack.Push("Three")
	stack.Push("Four")

	for s := range stack.Iterate() {
		if s == nil {
			break
		}
		fmt.Printf("%+v\n", s)
		assert.NotNil(t, s, "Nil")
	}
}

func TestStackLinkedList(t *testing.T) {
	stack := NewStackLinkedList()
	stack.Push("One")
	stack.Push("Two")
	stack.Push("Three")

	assert.Equal(t, 3, stack.Size(), "Not equal")

	item := stack.Pop()
	assert.Equal(t, "Three", item, "Not equal")

	stack.Push("Three")
	stack.Push("Four")

	for s := range stack.Iterate() {
		if s == nil {
			break
		}
		fmt.Printf("%+v\n", s)
		assert.NotNil(t, s, "Nil")
	}
}
