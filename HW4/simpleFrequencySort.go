package main

import (
	"fmt"
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

func main() {
	fmt.Println(Solve([]int{1, 2, 3, 4}))
	fmt.Println(Solve([]int{2, 3, 5, 3, 7, 9, 5, 3, 7}))
	fmt.Println(Solve([]int{1, 2, 3, 0, 5, 0, 1, 6, 8, 8, 6, 9, 1}))
}
