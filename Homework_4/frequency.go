package main

import (
	"fmt"
	"sort"
)

func frequency(arr []int) []int {
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

func main() {

	fmt.Println("Task 2 :", frequency([]int{2, 2, 2, 3, 4, 3, 4, 5, 6, 7, 8}))
}
