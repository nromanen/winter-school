package main

import "fmt"

func SumOfOdd(num int) int {
	return num * num * num
}
func main() {
	var num int

	fmt.Println("Input number:")
	fmt.Scan(&num)

	fmt.Println("Sum =", SumOfOdd(num))

	fmt.Println("\nInput number:")
	fmt.Scan(&num)

	fmt.Println("Sum =", SumOfOdd(num))

	fmt.Println("\nInput number:")
	fmt.Scan(&num)

	fmt.Println("Sum =", SumOfOdd(num))
}
