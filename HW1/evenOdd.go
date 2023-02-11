package main

import "fmt"

func main() {
	var x int
	fmt.Println("Input number:")
	fmt.Scan(&x)
	if x%2 == 0 {
		fmt.Println("The number is even")
	} else {
		fmt.Println("The number is odd")
	}
}
