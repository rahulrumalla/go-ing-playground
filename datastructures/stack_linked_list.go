package datastructures

type stackLinkedList struct {
	Current *LinkedListNode
	Length  int
}

type LinkedListNode struct {
	Item interface{}
	Next *LinkedListNode
}

func NewStackLinkedList() Stack {
	return &stackLinkedList{}
}

func (s *stackLinkedList) Size() int {
	return s.Length
}

func (s *stackLinkedList) IsEmpty() bool {
	return s.Length == 0
}

func (s *stackLinkedList) Push(obj interface{}) {
	newNode := &LinkedListNode{
		Item: obj,
		Next: s.Current,
	}
	s.Current = newNode
	s.Length++
}

func (s *stackLinkedList) Pop() interface{} {
	if !s.IsEmpty() {
		item := s.Current.Item
		s.Current = s.Current.Next
		s.Length--
		return item
	}
	return nil
}

func (s *stackLinkedList) Iterate() <-chan interface{} {
	ch := make(chan interface{}, 0)
	go func() {
		for {
			if s.IsEmpty() {
				break
			}
			ch <- s.Pop()
		}
		close(ch)
	}()
	return ch
}
