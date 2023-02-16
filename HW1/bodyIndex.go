package main

import "fmt"

func main() {
	var mass float32
	var height float32

	fmt.Println("please, enter your height in cm: ")
	fmt.Scan(&height)
	height /= 100
	fmt.Println("please, enter your mass in kg: ")
	fmt.Scan(&mass)

	var bmi = (mass / (height * height))
	fmt.Print("Your body mass index is: ", bmi)

	switch { // Сравнивает case с command
	case bmi < 16.0:
		fmt.Println(" (Underweight (Severe thinness))")
	case bmi < 17:
		fmt.Println(" (Underweight (Moderate thinness))")
	case bmi < 18.5:
		fmt.Println(" (Underweight (Mild thinness))")
	case bmi < 25:
		fmt.Println(" (Normal range)")
	case bmi < 30:
		fmt.Println(" (Overweight (Pre-obese))")
	case bmi < 35:
		fmt.Println(" (Obese (Class I))")
	case bmi < 40:
		fmt.Println(" (Obese (Class II))")
	case bmi >= 40:
		fmt.Println(" (Obese (Class III))")
	default:
		fmt.Println("Пока не ясно, що з вами не так")
	}
}
