package main

import (
	"fmt"
)

func Add(num int) func(int) int {
	
	return func(x int) int {
		return x + num
	}
}

func main() {
	f := Add(13)

	res := f(6)

	fmt.Println(res)
}
