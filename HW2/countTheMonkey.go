package main

import "fmt"

func monkeyCount(n int) (monkeys []int) {
	for i := 1; i <= n; i++ {
		monkeys = append(monkeys, i)
	}
	return
}

func main() {
	fmt.Println(monkeyCount(3))
	fmt.Println(monkeyCount(10))
	fmt.Println(monkeyCount(0))
}
