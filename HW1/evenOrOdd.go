package main

import "fmt"

func main() {

	var number int
	fmt.Println("Please, enter the number?")
	fmt.Scanf("%d\n", &number)

	var Type string
	if number%2 == 0 {
		Type = "Even"
	} else {
		Type = "Odd"
	}

	fmt.Printf("Hello, %d is ", number)
	fmt.Println(Type)
}
