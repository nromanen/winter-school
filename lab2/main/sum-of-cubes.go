package main

import "fmt"

func sumCubes(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i * i * i
	}
	return sum
}

func main() {
	// expected - 9
	fmt.Println(sumCubes(2))
	// expected -36
	fmt.Println(sumCubes(3))
	// expected - 3025
	fmt.Println(sumCubes(10))
}
