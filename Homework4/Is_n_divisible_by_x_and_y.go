package main

import "fmt"

func isDivisibleBy(n, x, y int) bool {
	return n%x == 0 && n%y == 0
}

func main() {
	var n, x, y int
	fmt.Println("Enter a number:")
	fmt.Scan(&n)
	fmt.Println("Enter two other numbers:")
	fmt.Scan(&x, &y)

	if isDivisibleBy(n, x, y) {
		fmt.Printf("%d is divisible by %d and %d\n", n, x, y)
	} else {
		fmt.Printf("%d is not divisible by %d and %d\n", n, x, y)
	}
}
