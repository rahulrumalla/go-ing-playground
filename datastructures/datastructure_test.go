package datastructures

// import (
// 	"fmt"
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// )

// func TestAppendToTailOnSinglyLinkedList(t *testing.T) {
// 	n := &Node{Item: 1}
// 	n.appendToTail(2)
// 	n.appendToTail(3)
// 	n.appendToTail(4)
// 	for {
// 		if n.Next == nil {
// 			assert.Equal(t, 4, n.Item, "Not equal")
// 			break
// 		} else {
// 			n = n.Next
// 		}
// 	}
// }

// func TestDeleteNodeOnSinglyLinkedList(t *testing.T) {
// 	n := &Node{Item: 1}
// 	n.appendToTail(2)
// 	n.appendToTail(3)
// 	n.appendToTail(4)

// 	itemToDelete := 4
// 	n = n.deleteNode(itemToDelete)
// 	for {
// 		fmt.Printf("%d -> ", n.Item)
// 		assert.True(t, n.Item != itemToDelete, "Not deleted")

// 		if n.Next == nil {
// 			break
// 		}

// 		n = n.Next
// 	}
// }
