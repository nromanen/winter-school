package main

import "fmt"

func monkeyCount(n int) (monkeys []int) {
	for i := 1; i <= n; i++ {
		monkeys = append(monkeys, i)
	}
	return
}

func main() {
	// expected - [1, 2, 3]
	fmt.Println(monkeyCount(3))
	// expected - [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]
	fmt.Println(monkeyCount(10))
	// expected - []
	fmt.Println(monkeyCount(0))
}
