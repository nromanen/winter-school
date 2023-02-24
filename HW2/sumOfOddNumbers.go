package main

import "fmt"

func RowSumOddNumbers(n int) int {
	c := 1
	for i := 1; i < n; i++ {
		c += i
	}
	firstNumberInRow := 2*c - 1
	sum := 0
	for i := 0; i < n; i++ {
		sum += firstNumberInRow + 2*i
	}
	return sum
}

func main() {
	fmt.Println(RowSumOddNumbers(2))
	fmt.Println(RowSumOddNumbers(13))
	fmt.Println(RowSumOddNumbers(0))
}
