package main

import "sort"
import "fmt"

func Solve(a []int) []int {
	sol := map[int]int{}
	for _, v := range a {
		sol[v]++
	}
	sort.Slice(a, func(i, j int) bool {
		if sol[a[i]] == sol[a[j]] {
			return a[i] < a[j]
		}
		return sol[a[i]] > sol[a[j]]
	})
	return a
}

func main() {
	fmt.Println(Solve([]int{2, 3, 5, 3, 7, 9, 5, 3, 7}))
}
