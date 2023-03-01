package main

import "fmt"

func main() {
	var x int
	fmt.Print("Input number: ")
	fmt.Scan(&x)
	index := x * x * x
	fmt.Printf("Result: %v", index)
}