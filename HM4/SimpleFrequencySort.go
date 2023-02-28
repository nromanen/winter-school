package kata

// https://www.codewars.com/kata/5a8d2bf60025e9163c0000bc/train/go

import (
	"sort"
)

func Solve(arr []int) []int {
	set := make(map[int]int)
	for _, v := range arr {
		set[v]++
	}
	sort.SliceStable(arr, func(i, j int) bool {
		if set[arr[i]] == set[arr[j]] {
			return arr[i] < arr[j]
		}
		return set[arr[i]] > set[arr[j]]
	})
	return arr
}
