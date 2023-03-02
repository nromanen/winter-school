package main

import (
	"fmt"
	"sort"
)

func sortByFrequency(arr []int) []int {
	freqMap := make(map[int]int)
	for _, num := range arr {
		if _, ok := freqMap[num]; ok {
			freqMap[num]++
		} else {
			freqMap[num] = 1
		}
	}

	sortedFreq := make([]int, 0, len(freqMap))
	for num := range freqMap {
		sortedFreq = append(sortedFreq, num)
	}
	sort.Slice(sortedFreq, func(i, j int) bool {
		if freqMap[sortedFreq[i]] == freqMap[sortedFreq[j]] {
			return sortedFreq[i] < sortedFreq[j]
		}
		return freqMap[sortedFreq[i]] > freqMap[sortedFreq[j]]
	})

	result := make([]int, 0, len(arr))
	for _, num := range sortedFreq {
		for i := 0; i < freqMap[num]; i++ {
			result = append(result, num)
		}
	}
	return result
}

func main() {
	arr := []int{2, 3, 5, 3, 7, 9, 5, 3, 7}
	sortedArr := sortByFrequency(arr)
	fmt.Println(sortedArr)
}