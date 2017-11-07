package datastructures

// import "errors"

// type Node struct {
// 	Item int
// 	Next *Node
// }

// func (n *Node) setNext(next *Node) {
// 	n.Next = next
// }

// func (n *Node) appendToTail(newItem int) {
// 	newNode := &Node{Item: newItem}
// 	for {
// 		if n.Next != nil {
// 			n = n.Next
// 		} else {
// 			break
// 		}
// 	}
// 	n.Next = newNode
// }

// // 4 -> 8 -> 6 -> 13
// // del 6
// func (head *Node) deleteNode(item int) *Node {
// 	if head.Item == item {
// 		head = head.Next
// 		return head
// 	}

// 	curr := head

// 	for {
// 		if curr.Next == nil {
// 			break
// 		}

// 		if curr.Next.Item == item {
// 			curr.Next = curr.Next.Next
// 			break
// 		}

// 		curr = curr.Next
// 	}

// 	return head
// }

// type Stack struct {
// 	Items []int
// }

// func (s Stack) push(v int) {
// 	s.Items = append(s.Items, v)
// }
