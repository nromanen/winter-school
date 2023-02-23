package main

import "fmt"

func SumCubes(num int) int {
	var sum int
	sum = 0

	for i := 1; i <= num; i++ {
		sum += i * i * i
	}

	return sum
}
func main() {
	fmt.Println(SumCubes(3))

	fmt.Println(SumCubes(5))

}
