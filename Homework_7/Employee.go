package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
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
	file, err := os.Open("employees.csv")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV:", err)
		return
	}

	employees := make([]Employee, len(records))
	for i, record := range records {
		birthdate, err := time.Parse("2006-01-02", record[1])
		if err != nil {
			fmt.Println("Error parsing birthdate:", err)
			continue
		}

		salary, err := strconv.ParseFloat(record[2], 64)
		if err != nil {
			fmt.Println("Error parsing salary:", err)
			continue
		}

		employees[i] = Employee{
			FullName:  record[0],
			BirthDate: birthdate,
			Salary:    salary,
		}
	}

	for _, employee := range employees {
		fmt.Println("Employee Full Name:", employee.FullName)
		fmt.Println("Employee Birth Date:", employee.BirthDate.Format("02-Jan-2006"))
		fmt.Println("Employee Salary:", employee.Salary)
		fmt.Println("Employee Age:", employee.GetAge())
		fmt.Println("---------------------------")
	}
	file.Close()
}
