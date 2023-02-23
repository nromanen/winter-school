package main

import "fmt"

func IsDivisible(n, x, y int) bool {
	return n%x == 0 && n%y == 0
}

func main() {
	fmt.Println(IsDivisible(3, 1, 3))
	fmt.Println(IsDivisible(12, 2, 6))
	fmt.Println(IsDivisible(100, 5, 3))
	fmt.Println(IsDivisible(12, 7, 5))
	fmt.Println(IsDivisible(3, 1, 2))
}
