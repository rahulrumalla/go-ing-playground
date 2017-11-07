package leetcode

import "sort"

type ByRune []rune

func (r ByRune) Less(i, j int) bool {
	return r[i] < r[j]
}

func (r ByRune) Len() int {
	return len(r)
}

func (r ByRune) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

// ["eat", "tea", "tan", "ate", "nat", "bat"]
func GroupAnagrams(strs []string) [][]string {
	sortedStrs := make(map[string][]string, len(strs))
	for _, str := range strs {
		r := []rune(str)
		sort.Sort(ByRune(r))
		sr := string(r)
		sortedStrs[sr] = append(sortedStrs[sr], str)
	}

	var output [][]string
	for _, list := range sortedStrs {
		output = append(output, list)
	}

	return output
}
