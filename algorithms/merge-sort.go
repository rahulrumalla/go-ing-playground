package algorithms

func MergeSort(items []int) []int {
	length := len(items)

	if length < 2 {
		return items
	}

	mid := length / 2

	return Merge(MergeSort(items[:mid]), MergeSort(items[mid:]))
}

func Merge(left, right []int) []int {
	size := len(left) + len(right)
	i, j := 0, 0

	mergedSlice := make([]int, size, size)

	for k := 0; k < size; k++ {
		if i > len(left)-1 && j <= len(right)-1 {
			// left was exhausted
			mergedSlice[k] = right[j]
			j++
		} else if i <= len(left)-1 && j > len(right)-1 {
			// right was exhausted
			mergedSlice[k] = left[i]
			i++
		} else if left[i] <= right[j] {
			mergedSlice[k] = left[i]
			i++
		} else {
			mergedSlice[k] = right[j]
			j++
		}
	}

	return mergedSlice
}
