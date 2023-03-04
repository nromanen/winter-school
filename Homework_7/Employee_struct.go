package main

import (
	"fmt"
)

type Employee struct {
	FullName  string
	BirthDate string
	Salary    float64
}

func main() {
	employee := Employee{
		FullName:  "Nick Bush",
		BirthDate: "10.05.2000",
		Salary:    50000.00,
	}

	fmt.Println("Employee Full Name:", employee.FullName)
	fmt.Println("Employee Birth Date:", employee.BirthDate)
	fmt.Println("Employee Salary:", employee.Salary)
}
