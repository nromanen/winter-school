package main

import "fmt"

func main() {
	var x int
	fmt.Println("Input number:")
	fmt.Scan(&x)
	if x%2 == 0 {
		fmt.Printf("%v is even", x)
	} else {
		fmt.Printf("%v is odd", x)
	}
}
