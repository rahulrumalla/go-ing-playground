package leetcode

// [3, 4, 6, 2, 1, 7] 8
// [2, 3]

func TwoSum(nums []int, target int) []int {
	keyMap := make(map[int]int, len(nums))

	for key, num := range nums {
		if km, exists := keyMap[target-num]; exists {
			return []int{km, key}
		}

		keyMap[num] = key
	}
	return []int{}
}
