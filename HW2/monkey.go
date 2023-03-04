package main

import "fmt"

func monkeyCount(n int) []int {
	count := []int{}

	for i := 0; i < n; i++ {
		count = append(count, i+1)
	}

	return count
}
func main() {
	fmt.Println("for example, there are 10 monkeys:", monkeyCount(10))

}
