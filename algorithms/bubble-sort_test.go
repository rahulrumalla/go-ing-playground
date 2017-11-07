package algorithms

import (
	"fmt"
	"testing"
)

var (
	input = []int{21, 4, 2, 13, 10, 0, 19, 11, 7, 5, 23, 18, 9, 14, 6, 8, 1, 20, 17, 3, 16, 22, 24, 15, 12}
)

func TestBubbleSort(t *testing.T) {
	fmt.Printf("Unsorted: %+v\n", input)
	BubbleSort(input, false)
	//sort.Ints(input)
	fmt.Printf("Sorted: %+v\n", input)
}

func TestMergeSort(t *testing.T) {
	fmt.Printf("Unsorted: %+v\n", input)
	ouput := MergeSort(input)
	fmt.Printf("Sorted: %+v\n", ouput)
}

func TestQuickSort(t *testing.T) {
	fmt.Printf("Unsorted: %+v\n", input)
	ouput := QuickSort(input)
	fmt.Printf("Sorted: %+v\n", ouput)
}
