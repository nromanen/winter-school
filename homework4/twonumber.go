package main

import "fmt"

func isDivisibleByTwoNumbers(n, x, y int) bool {
	if n%x == 0 && n%y == 0 {
		return true
	}
	return false
}

func main() {

	fmt.Println(isDivisibleByTwoNumbers(3, 1, 3))   // true
	fmt.Println(isDivisibleByTwoNumbers(12, 2, 6))  // true
	fmt.Println(isDivisibleByTwoNumbers(100, 5, 3)) // false
	fmt.Println(isDivisibleByTwoNumbers(12, 7, 5))  // false

}
