package main

import "fmt"

func main() {

	var num int
	fmt.Println("Input number:")
	fmt.Scan(&num)
	if num%2 == 0 {
		fmt.Println(num, "is even number")
	} else {
		fmt.Println(num, "is odd number")
	}
}
