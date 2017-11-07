package datastructures

type stackArray struct {
	Items []interface{}
	Index int
}

func NewStackArray() Stack {
	s := &stackArray{}
	s.Items = make([]interface{}, 0)
	return s
}

func (s *stackArray) Push(obj interface{}) {
	s.Items = append(s.Items, obj)
	s.Index++

	// array cap needs to be resized
}

func (s *stackArray) Pop() interface{} {
	if s.Index == 0 {
		return nil
	}

	item := s.Items[s.Index-1]
	s.Items = s.Items[0 : s.Index-1]

	// array cap needs to be resized

	s.Index--
	return item
}

func (s *stackArray) IsEmpty() bool {
	return s.Size() == 0
}

func (s *stackArray) Size() int {
	return len(s.Items)
}

func (s *stackArray) Iterate() <-chan interface{} {
	ch := make(chan interface{})
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
