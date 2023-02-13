package main

import "fmt"

func main() {
	var m int
	var h float64

	fmt.Println("Input your body weight:")
	fmt.Scan(&m)

	fmt.Println("Input your height(in meters):")
	fmt.Scan(&h)

	index := float64(m) / (h * h)

	fmt.Println("Your BMI category:")
	switch {
	case index < 16.0:
		fmt.Printf("BMI:%.2f - Underweight (Severe thinness)", index)
	case index < 17.0:
		fmt.Printf("BMI:%.2f - Underweight (Moderate thinness)", index)
	case index < 18.5:
		fmt.Printf("BMI:%.2f - Underweight (Mild thinness)", index)
	case index < 25.0:
		fmt.Printf("BMI:%.2f - Normal range", index)
	case index < 30.0:
		fmt.Printf("BMI:%.2f - Overweight (Pre-obese)", index)
	case index < 35.0:
		fmt.Printf("BMI:%.2f - Obese (Class I)", index)
	case index < 40.0:
		fmt.Printf("BMI:%.2f - Obese (Class II)", index)
	case index >= 40.0:
		fmt.Printf("BMI:%.2f - Obese (Class III)", index)
	default:
		fmt.Printf("Error")
	}
}
