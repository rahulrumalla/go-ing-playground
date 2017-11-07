package datastructures

type qLinkedList struct {
	Head   *LinkedListNode
	Tail   *LinkedListNode
	Length int
}

func NewQLinkedList() Queue {
	return &qLinkedList{}
}

func (q *qLinkedList) Size() int {
	return q.Length
}

func (q *qLinkedList) IsEmpty() bool {
	return q.Length == 0
}

func (q *qLinkedList) Enqueue(obj interface{}) {
	oldTail := q.Tail

	newNode := &LinkedListNode{
		Item: obj,
	}
	q.Tail = newNode

	if q.IsEmpty() {
		q.Head = q.Tail
	} else {
		oldTail.Next = q.Tail
	}

	q.Length++
}

func (q *qLinkedList) Dequeue() interface{} {
	var item interface{}
	if !q.IsEmpty() {
		item = q.Head.Item
		q.Head = q.Head.Next
		q.Length--
		if q.Length == 0 {
			q.Tail = q.Head
		}
	}
	return item
}

func (q *qLinkedList) Iterate() <-chan interface{} {
	ch := make(chan interface{}, 0)
	go func() {
		for {
			if q.IsEmpty() {
				break
			}
			ch <- q.Dequeue()
		}
		close(ch)
	}()
	return ch
}
