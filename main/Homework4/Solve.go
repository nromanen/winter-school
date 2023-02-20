package kata

import "sort"

//Task 2 https://www.codewars.com/kata/5a8d2bf60025e9163c0000bc
func Solve(arr []int) []int {
	m := make(map[int]int, 0)

	for _, n := range arr {
		m[n]++
	}

	sort.Slice(arr, func(i, j int) bool {
		if m[arr[i]] == m[arr[j]] {
			return arr[i] < arr[j]
		}

		return m[arr[i]] > m[arr[j]]
	})

	return arr
}