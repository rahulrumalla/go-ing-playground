package algorithms

import (
	"fmt"
)

func BubbleSort(numbers []int, ascending bool) {
	N := len(numbers)
	var i int
	for i = 0; i < N; i++ {
		if !sweep(numbers, i, ascending) {
			return
		}
	}
}

func sweep(numbers []int, prevPasses int, ascending bool) bool {
	N := len(numbers)

	fi := 0
	si := 1

	swapped := false

	for si < (N - prevPasses) {
		fn := numbers[fi]
		sn := numbers[si]

		if (ascending && fn > sn) || (!ascending && fn < sn) {
			numbers[fi] = sn
			numbers[si] = fn
			swapped = true
		}

		fi++
		si++
	}
	fmt.Printf("Sweep result: %+v\n", numbers)

	return swapped
}
