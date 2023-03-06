package main

import (
	"fmt"
)

func main() {
	var number int
	fmt.Print("Enter a number: ")
	fmt.Scan(&number)

	if number%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}

	var month int
	fmt.Print("Enter a month (1-12): ")
	fmt.Scan(&month)

	quarter := (month-1)/3 + 1
	fmt.Println("Quarter:", quarter)

	var m int
	var h float64
	fmt.Print("Enter weight (kg): ")
	fmt.Scan(&m)
	fmt.Print("Enter height (m): ")
	fmt.Scan(&h)

	bmi := float64(m) / (h * h)

	if bmi < 18.5 {
		fmt.Println("Underweight")
	} else if bmi < 25 {
		fmt.Println("Normal")
	} else if bmi < 30 {
		fmt.Println("Overweight")
	} else {
		fmt.Println("Obese")
	}
}
