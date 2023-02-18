package main

import "fmt"

func IsDivisible(n, x, y int) bool {
	return n%x == 0 && n%y == 0
}

func main() {
	// expected - true
	fmt.Println(IsDivisible(3, 1, 3))
	// expected - true
	fmt.Println(IsDivisible(12, 2, 6))
	// expected - false
	fmt.Println(IsDivisible(100, 5, 3))
	// expected - false
	fmt.Println(IsDivisible(12, 7, 5))
}
