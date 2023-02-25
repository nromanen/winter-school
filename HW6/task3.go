package main

import "fmt"

func main() {
	f := Add(5)
	fmt.Println(f(10)) // Output: 15
	fmt.Println(f(20)) // Output: 25
}

func Add(n int) func(int) int {
	return func(x int) int {
		return x + n
	}
}
