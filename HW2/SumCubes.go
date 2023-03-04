package main

import "fmt"

func SumCubes(n int) int {
	var sum int

	for i := 0; i < n; n-- {

		sum = (n * n * n) + sum
	}
	return sum

}

func main() {
	fmt.Println("for example, SumCubes of 5:", SumCubes(5))
}
