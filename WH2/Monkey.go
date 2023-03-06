package main

import "fmt"

func main() {
	var monkey []int
	var x int
	fmt.Print("Input number: ")
	fmt.Scan(&x)
	for i := 0; i <= x; i++ {
		monkey = append(monkey, i)
	}
	fmt.Print(monkey)
}
