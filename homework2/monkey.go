package main

import "fmt"

func countMonkeys(n int) []int {
	nums := make([]int, n)
	for i := 1; i <= n; i++ {
		nums[i-1] = i
	}
	return nums
}

func main() {
	fmt.Println(countMonkeys(10)) // [1 2 3 4 5 6 7 8 9 10]
	fmt.Println(countMonkeys(1))  // [1]
	fmt.Println(countMonkeys(0))  // []
}
