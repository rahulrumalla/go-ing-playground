package datastructures

type Stack interface {
	Push(obj interface{})
	Pop() interface{}
	IsEmpty() bool
	Size() int
	Iterate() <-chan interface{}
}
