package kata

import "sort"

func Solve(a []int) []int {
	result := map[int]int{}
	for _, v := range a {
		result[v]++
	}
	sort.Slice(a, func(i, j int) bool {
		if result[a[i]] == result[a[j]] {
			return a[i] < a[j]
		}
		return result[a[i]] > result[a[j]]
	})
	return a
}
