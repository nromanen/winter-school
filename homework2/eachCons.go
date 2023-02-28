package main

import "fmt"

func eachCons(lst []int, n int) [][]int {
	result := [][]int{}
	for i := 0; i <= len(lst)-n; i++ {
		result = append(result, lst[i:i+n])
	}
	return result
}

func main() {
	fmt.Println(eachCons([]int{1, 2, 3, 4}, 2)) // Output: [[1 2] [2 3] [3 4]]
	fmt.Println(eachCons([]int{1, 2, 3, 4}, 3)) // Output: [[1 2 3] [2 3 4]]
}