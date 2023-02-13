package main

import "fmt"

func main() {
	var mass float64
	var height float64
	fmt.Print("Введіть вашу вагу в кг: ")
	fmt.Scan(&mass)

	for {
		if mass > 0 {
			break
		}
		fmt.Print("Маса не може бути менше нуля. Введіть масу в кг ще раз: ")
		fmt.Scan(&mass)
	}

	fmt.Print("Введіть ваш ріст в метрах: ")
	fmt.Scan(&height)

	for {
		if height > 0 {
			break
		}
		fmt.Print("Ріст не може бути менше нуля. Введіть ріст в метрах ще раз: ")
		fmt.Scan(&height)
	}

	var index float64 = mass / (height * height)
	fmt.Printf("Ваш індекс маси тіла = %g. Ваша класифікація:\n", index)
	if index < 18.5 {
		fmt.Println("Недостатня маса")
	} else if index >= 18.5 && index < 25 {
		fmt.Println("Норма")
	} else if index >= 25 && index < 30 {
		fmt.Println("Переожиріння(гладкість)")
	} else if index >= 30 && index < 35 {
		fmt.Println("Ожиріння 1 ступення")
	} else if index >= 35 && index < 40 {
		fmt.Println("Ожиріння 2 ступення")
	} else {
		fmt.Println("Ожиріння 3 ступення")
	}
}
