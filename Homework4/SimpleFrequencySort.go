package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{2, 3, 5, 3, 7, 9, 5, 3, 7}
	freqMap := make(map[int]int)
	for _, num := range arr {
		freqMap[num]++
		fmt.Println(freqMap)
	}

	sort.Slice(arr, func(i, j int) bool {
		if freqMap[arr[i]] != freqMap[arr[j]] {
			return freqMap[arr[i]] > freqMap[arr[j]]
		}
		return arr[i] < arr[j]
	})

	fmt.Println(arr)
}
