package main

import "fmt"

func main() {
	var number int
	fmt.Print("Enter your number: ")
	fmt.Scan(&number)
	if number%2 == 0 {
		fmt.Println("It's an even number")
	} else {
		fmt.Println("It's an odd number")
	}
}
