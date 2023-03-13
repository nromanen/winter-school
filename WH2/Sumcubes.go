package main

import (
	"fmt"
)

func sumCubes(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i * i * i
	}
	return sum
}

func main() {
	fmt.Println(sumCubes(3)) // Output: 36
	fmt.Println(sumCubes(5)) // Output: 225
}
