package main

import (
	"fmt"
)

func main() {
	var m, h float64
	fmt.Print("Enter the mass (kg): ")
	fmt.Scan(&m)
	fmt.Print("Enter the height (m): ")
	fmt.Scan(&h)

	bmi := m / (h * h)
	if bmi < 18.5 {
		fmt.Println("Underweight")
	} else if bmi >= 18.5 && bmi < 25 {
		fmt.Println("Normal weight")
	} else if bmi >= 25 && bmi < 30 {
		fmt.Println("Overweight")
	} else if bmi >= 30 && bmi < 35 {
		fmt.Println("Obesity Level I")
	} else if bmi >= 35 && bmi < 40 {
		fmt.Println("Obesity Level II")
	} else if bmi >= 40 {
		fmt.Println("Obesity Level III")
	}
}
