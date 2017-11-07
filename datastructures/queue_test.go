package datastructures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueueLinkedList(t *testing.T) {
	q := NewQLinkedList()
	q.Enqueue("One")
	q.Enqueue("Two")
	q.Enqueue("Three")

	assert.Equal(t, 3, q.Size(), "Not true")

	val := q.Dequeue()
	assert.Equal(t, "One", val, "Not equal")

	val = q.Dequeue()
	assert.Equal(t, "Two", val, "Not equal")

	val = q.Dequeue()
	assert.Equal(t, "Three", val, "Not equal")

	val = q.Dequeue()
	assert.Nil(t, val, "Not Nil")
}
