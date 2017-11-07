package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(x *ListNode, y *ListNode) *ListNode {
	res := &ListNode{}

	// x := l1
	// y := r1

	carry := 0
	curr := res
	for x != nil || y != nil {
		xVal, yVal := 0, 0
		if x != nil {
			xVal = x.Val
		}
		if y != nil {
			yVal = y.Val
		}

		// fmt.Printf("%d | %d\n", xVal, yVal)

		sum := xVal + yVal + carry

		if sum >= 10 {
			carry = 1
			sum -= 10

			// fmt.Printf("Sum: %d | Carry: %d\n", sum, carry)
		} else {
			carry = 0
		}

		curr.Next = &ListNode{Val: sum} // reflects back to res
		curr = curr.Next                // move curr obj position

		if x != nil {
			x = x.Next
		}

		if y != nil {
			y = y.Next
		}
	}

	if carry > 0 {
		curr.Next = &ListNode{Val: carry}
	}

	return res.Next
}
