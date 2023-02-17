package main

import "fmt"

func OddValue(n int) int {
	return n*n*n
}

func main() {
	fmt.Println(OddValue(2),OddValue(13),OddValue(0))
	
}