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
	// expected - [1 2 3 4]
	fmt.Println(Solve([]int{1, 2, 3, 4}))
	// expected - [3 3 3 5 5 7 7 2 9]
	fmt.Println(Solve([]int{2, 3, 5, 3, 7, 9, 5, 3, 7}))
	// expected - [1 1 1 0 0 6 6 8 8 2 3 5 9]
	fmt.Println(Solve([]int{1, 2, 3, 0, 5, 0, 1, 6, 8, 8, 6, 9, 1}))
}
