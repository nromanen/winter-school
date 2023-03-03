package main

import "fmt"

func Add(n int) func(int) int {
	return func(x int) int {
		return x + n
	}
}

func main() {
	f := Add(5)
	fmt.Println(f(10))
}
