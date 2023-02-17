package main

import ("fmt")

func main() {
	var mass uint
	var height float64
	var index float64

	fmt.Println("Enter mass:")
	fmt.Scan(&mass)
	fmt.Println("And your heigth:")
	fmt.Scan(&height)

	index = float64(mass) / (height * height)
	fmt.Println("BMI category:")

	switch {
	case index < 18.5:
		fmt.Println("Insufficient mass")
	case index >= 18.5 && index < 25:
		fmt.Println("Norm")
	case index >= 25 && index < 30:
		fmt.Println("Obesity")
	case index >= 30 && index < 35:
		fmt.Println("Obesity of the I degree")
	case index >= 35 && index < 40:
		fmt.Println("Obesity of the II degree")
	case index >= 40:
		fmt.Println("Obesity of the III degree")

	default:
		fmt.Println("Err")

	}

}