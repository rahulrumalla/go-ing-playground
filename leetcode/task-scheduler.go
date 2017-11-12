package leetcode

import "sort"

func LeastInterval(tasks []byte, n int) int {
	taskMap := make([]int, 26)

	for _, t := range tasks {
		taskMap[t-'A']++
	}

	sort.Sort(sort.Reverse(sort.IntSlice(taskMap)))

	return 0
}

func TestLeastInterval() {
	LeastInterval([]byte{'A', 'A', 'B', 'B', 'B', 'C'}, 2)
}
