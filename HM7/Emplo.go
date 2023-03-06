package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"
)

type Employee struct {
	FullName   string
	BirthDate  string
	Salary     float64
	Department string
}

func (e *Employee) GetAge() int {
	now := time.Now()
	birthDate, err := time.Parse("15.12.2001", e.BirthDate)
	if err != nil {
		fmt.Printf("Error parsing birth date for %s: %s\n", e.FullName, err)
		return 0
	}
	ageDuration := now.Sub(birthDate)
	ageInYears := int(ageDuration.Hours() / 24 / 365)
	return ageInYears
}

func averageSalaryByAgeGroup(employees []Employee) map[int]float64 {
	salaryByAge := make(map[int][]float64)
	for _, employee := range employees {
		age := employee.GetAge()
		salaryByAge[age] = append(salaryByAge[age], employee.Salary)
	}

	avgSalaryByAgeGroup := make(map[int]float64)
	for age, salaries := range salaryByAge {
		ageGroup := (age-1)/10 + 1
		totalSalary := 0.0
		for _, salary := range salaries {
			totalSalary += salary
		}
		avgSalaryByAgeGroup[ageGroup] = totalSalary / float64(len(salaries))
	}

	return avgSalaryByAgeGroup
}

func main() {
	file, err := os.Open("emplo.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Create a CSV
	reader := csv.NewReader(file)
	reader.Comma = ';'
	reader.FieldsPerRecord = 4

	// Read
	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	employees := make([]Employee, len(records)-1)

	for i, record := range records[1:] {
		fullName := record[0]
		birthDate := record[1]
		salary, _ := strconv.ParseFloat(record[2], 64)
		department := record[3]

		employee := Employee{
			FullName:   fullName,
			BirthDate:  birthDate,
			Salary:     salary,
			Department: department,
		}

		employees[i] = employee
	}

	for _, employee := range employees {
		fmt.Printf("%s is %d years old, earns %.2f and is in %s department\n", employee.FullName, employee.GetAge(), employee.Salary)
	}

	avgSalaryByAgeGroup := averageSalaryByAgeGroup(employees)
	for ageGroup, avgSalary := range avgSalaryByAgeGroup {
		fmt.Printf("Average salary for age group %d is %.2f\n", ageGroup, avgSalary)
	}
}
