package kata

import "sort"

func Solve(arr []int) []int {

	sortCount := make(map[int]int)

	for _, v := range arr {
		sortCount[v]++
	}

	sort.Slice(arr, func(i, j int) bool {
		if sortCount[arr[i]] == sortCount[arr[j]] {
			return arr[i] < arr[j]
		}
		return sortCount[arr[i]] > sortCount[arr[j]]
	})
	return arr

	//    solve([2,3,5,3,7,9,5,3,7]) = [3,3,3,5,5,7,7,2,9]
	//    sort by highest to lowest
}
