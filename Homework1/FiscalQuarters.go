package main

import "fmt"

func main() {

	var month int
	fmt.Println("Enter month in this format 1_12:")
	fmt.Scanf("%d\n", &month)

	if month <= 3 {
		fmt.Println("it's 1")
	} else if month <= 6 {
		fmt.Println("it's 2")
	} else if month <= 9 {
		fmt.Println("it's 3")
	} else {
		fmt.Println("it's 4")
	}
	fmt.Println("quarter")

}