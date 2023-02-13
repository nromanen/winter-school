package main

import "fmt"

func main() {
	var number int
	fmt.Println("Input number : ")
	fmt.Scan(&number)
	if number%2 == 0 {
		fmt.Printf("%d is a Even number", number)
	} else {
		fmt.Printf("%d is a Odd number", number)
	}

}