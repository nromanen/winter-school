package main

import (
	"fmt"
	"time"
)

type Employee struct {
	Name        string
	DateOfBirth time.Time

}

func (e *Employee) GetAge() int {
	today := time.Now()
	age := today.Year() - e.DateOfBirth.Year()


	if today.Month() < e.DateOfBirth.Month() ||
		(today.Month() == e.DateOfBirth.Month() && today.Day() < e.DateOfBirth.Day()) {
		age -= 1
	}

	return age
}

func main() {

	employee := Employee{
		Name:        "John",
		DateOfBirth: time.Date(1980, time.January, 1, 0, 0, 0, 0, time.UTC),

	}

	age := employee.GetAge()

	fmt.Printf("%s is %d years old\n", employee.Name, age)

}
