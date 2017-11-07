package leetcode

import (
	"fmt"
	"sort"
)

// [2, 2, 3, 4, 6, 7, 4]

func TriangleNumber(nums []int) int {
	count := 0

	sort.Ints(nums)

	iter := 0

	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			for k := j + 1; k < len(nums); k++ {
				if isTraingle(nums[i], nums[j], nums[k]) {
					count++
					// fmt.Printf("%d %d %d\n", nums[i], nums[j], nums[k])
				} else {
					break
				}
				iter++
			}
		}
	}

	fmt.Printf("Iterations: %d", iter)

	return count
}

func isTraingle(a, b, c int) bool {
	return a+b > c //&& b+c > a && a+c > b // because a < b < c
}
