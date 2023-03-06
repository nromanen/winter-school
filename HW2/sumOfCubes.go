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
	fmt.Println(sumCubes(2))
	fmt.Println(sumCubes(3))
	fmt.Println(sumCubes(10))
}
