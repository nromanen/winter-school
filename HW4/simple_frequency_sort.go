// https://www.codewars.com/kata/5a8d2bf60025e9163c0000bc/train/go

package kata

import "sort"

func Solve(arr []int) []int {
	// Count the frequency of each element in the array
	freqMap := make(map[int]int)
	for _, val := range arr {
		freqMap[val]++
	}

	// Create a slice of pairs (element, frequency) from the map
	pairs := make([][2]int, 0, len(freqMap))
	for k, v := range freqMap {
		pairs = append(pairs, [2]int{k, v})
	}

	// Sort the slice of pairs by decreasing frequency
	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i][1] != pairs[j][1] {
			return pairs[i][1] > pairs[j][1]
		}
		return pairs[i][0] < pairs[j][0]
	})

	// Create a sorted array based on the sorted pairs
	sortedArr := make([]int, len(arr))
	i := 0
	for _, pair := range pairs {
		for j := 0; j < pair[1]; j++ {
			sortedArr[i] = pair[0]
			i++
		}
	}

	return sortedArr
}
