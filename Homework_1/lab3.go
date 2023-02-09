package main

import "fmt"

func main() {
	var (
		m int
		h float32
	)
	fmt.Println("Input your body mass:")
	fmt.Scan(&m)
	fmt.Println("Input your height (meters):")
	fmt.Scan(&h)

	var ind = float32(m) / (h * h)
	//fmt.Println(ind)
	fmt.Println("Your BMI category:")
	switch {
	case ind > 0 && ind <= 16.0:
		fmt.Println("Underweight (Severe thinness)")
	case ind > 16.0 && ind <= 16.9:
		fmt.Println("Underweight (Moderate thinness)")
	case ind >= 17.0 && ind <= 18.4:
		fmt.Println("Underweight (Mild thinness)")
	case ind >= 18.5 && ind <= 24.9:
		fmt.Println("Normal range")
	case ind >= 25.0 && ind <= 29.9:
		fmt.Println("Overweight (Pre-obese)")
	case ind >= 30.0 && ind <= 34.9:
		fmt.Println("Obese (Class I)")
	case ind >= 35.0 && ind <= 39.9:
		fmt.Println("Obese (Class II)")
	case ind >= 40.0:
		fmt.Println("Obese (Class III)")
	default:
		fmt.Println("Invalid input")
	}

}
