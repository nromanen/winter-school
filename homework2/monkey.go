package main

import "fmt"

func monkey(n int) []int {
	var array []int
	for i := 1; i <= n; i++ {
		array = append(array, i)
	}
	return array

}

func main() {

	fmt.Println(monkey(10))
	fmt.Println(monkey(1))
	fmt.Println(monkey(0))

}