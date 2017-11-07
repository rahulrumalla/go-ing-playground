package algorithms

import "math/rand"

func QuickSort(items []int) []int {
	// handle 0 & 1 size list
	if len(items) < 2 {
		return items
	}

	// set left & right
	l, r := 0, len(items)-1

	// set random pivot index to split the list ( l <-> pi <-> r)
	pi := rand.Int() % len(items)

	// swap r index with pivot index
	items[pi], items[r] = items[r], items[pi]

	// iterate from i=0 to r.
	for i := range items {
		// Swap item at i with item at l if item at i < item at r
		if items[i] < items[r] {
			items[i], items[l] = items[l], items[i]

			// move l index ahead
			l++
		}
	}

	items[l], items[r] = items[r], items[l]

	// recursive to list from 0 to l
	QuickSort(items[:l])

	// recursive to list from l+1 to the end
	QuickSort(items[l+1:])

	return items
}

// 22, 3, 4, 76, 24, 56, 2, 1, 12
// 22, 3, 12, 76, 24, 56, 2, 1, 4
// 3, 2, 12, 76, 24, 56, 22, 1, 4
// 3, 2, 1, 76, 24, 56, 22, 12, 4
// 3, 2, 1, 4, 24, 56, 22, 12, 76
