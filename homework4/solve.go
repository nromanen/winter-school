package main

import (
	"fmt"
	"sort"
)

func solve(arr []int) []int {
	// Count the frequency of each element
	freq := make(map[int]int)
	for _, num := range arr {
		freq[num]++
	}

	// Create a slice of key-value pairs
	pairs := make([][2]int, 0, len(freq))
	for k, v := range freq {
		pairs = append(pairs, [2]int{k, v})
	}

	// Sort the key-value pairs by decreasing frequency
	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i][1] == pairs[j][1] {
			return pairs[i][0] < pairs[j][0]
		}
		return pairs[i][1] > pairs[j][1]
	})

	// Create a sorted slice of integers
	sorted := make([]int, 0, len(arr))
	for _, pair := range pairs {
		for i := 0; i < pair[1]; i++ {
			sorted = append(sorted, pair[0])
		}
	}

	return sorted
}

func main() {

	arr := []int{2, 3, 5, 3, 7, 9, 5, 3, 7}
	fmt.Println(solve(arr)) // [3, 3, 3, 5, 5, 7, 7, 2, 9]

}
