package main

import "fmt"

func main() {
	var x int
	fmt.Println("Input number:")
	fmt.Scan(&x)
	if x%2 == 0 {
		fmt.Println("x is even")
	} else {
		fmt.Println("x is odd")
	}
}
