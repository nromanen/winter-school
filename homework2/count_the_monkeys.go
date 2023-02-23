package main

import "fmt"

func CountMonkeys(num int) []int {
	count := make([]int, num)

	for i := 0; i < num; i++ {
		count[i] = i + 1
	}
	return count
}

func main() {
	var num int
	fmt.Scan(&num)
	fmt.Println(CountMonkeys(num))
}
