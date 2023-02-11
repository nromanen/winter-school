package main

import "fmt"

func main() {
	var month int
	fmt.Print("Enter a month as an integer from 1 to 12: ")
	fmt.Scan(&month)

	quarter := (month-1)/3 + 1
	fmt.Printf("The month belongs to the %d quarter of the year.\n", quarter)
}
