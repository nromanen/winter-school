package main

import "fmt"

//Task 3

func Add(n int) func(int) int {
	return func(x int) int {
		return x + n
	}
}

func main() {
	addFive := Add(5)
	result := addFive(10)
	fmt.Println(result) // Output: 15

	addTen := Add(10)
	result = addTen(20)
	fmt.Println(result) // Output: 30
}
