package main

import (
	"fmt"
	"time"
)

type Employee struct {
	FullName  string
	BirthDate time.Time
	Salary    float64
}

func (e Employee) GetAge() int {
	today := time.Now()
	age := today.Year() - e.BirthDate.Year()
	if today.YearDay() < e.BirthDate.YearDay() {
		age--
	}
	return age
}

func main() {
	employee := Employee{
		FullName:  "Nick Bush",
		BirthDate: time.Date(2000, time.March, 12, 0, 0, 0, 0, time.UTC),
		Salary:    50000.00,
	}

	fmt.Println("Employee Full Name:", employee.FullName)
	fmt.Println("Employee Birth Date:", employee.BirthDate.Format("02-Jan-2006"))
	fmt.Println("Employee Salary:", employee.Salary)
	fmt.Println("Employee Age:", employee.GetAge())
}
