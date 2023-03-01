package main

import "fmt"

func sum_of_cubes(x int) int {
	var total int = 0
	for x > 0 {
		total += x * x * x
		x -= 1
	}
	return total
}

func main() {
	var x int
	fmt.Print("Input number: ")
	fmt.Scan(&x)

	fmt.Print(sum_of_cubes(x))
}