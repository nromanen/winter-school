package main

import "fmt"

func main() {
	var number int
	fmt.Scanf("%d\n", &number)
	if number%2 == 0 {
		fmt.Printf("Even")
	} else {
		fmt.Printf("Odd")
	}
}
