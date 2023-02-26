package main

import (
	"fmt"
)

func main() {
	f := Add(25)
	result := f(10)
	fmt.Println(result) //will be 35
}

func Add(n int) func(int) int {
	return func(x int) int {
		return x + n
	}
}
