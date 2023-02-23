package main

import "fmt"

func EachCons(num []int, n int) [][]int {

	var subset [][]int

	for i := 0; n <= len(num); i++ {
		subset = append(subset, num[i:n])
		n++
	}

	return subset
}

func main() {

	fmt.Println(EachCons([]int{1, 2, 3, 4}, 2))

	fmt.Println(EachCons([]int{1, 2, 3, 4, 5, 6, 7}, 5))

	fmt.Println(EachCons([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 8))
}
