package main

import (
	"fmt"
)

func IsDivisible(n, x, y int) bool {
	if (n%x)+(n%y) == 0 {
		return true
	}

	return false
}

func main() {

	fmt.Println("Task 1 :", IsDivisible(10, 2, 5))
}
