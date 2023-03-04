package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"
)

// Define a struct to hold employee information
type Employee struct {
	FullName   string
	BirthDate  string
	Salary     float64
	Department string
}

// Define a method to calculate employee age based on birth date
func (e *Employee) GetAge() int {
	now := time.Now()
	birthDate, err := time.Parse("02.01.2006", e.BirthDate)
	if err != nil {
		fmt.Printf("Error parsing birth date for %s: %s\n", e.FullName, err)
		return 0
	}
	ageDuration := now.Sub(birthDate)
	ageInYears := int(ageDuration.Hours() / 24 / 365)
	return ageInYears
}

// Define a function to calculate the average salary by age group for a list of employees
func averageSalaryByAgeGroup(employees []Employee) map[int]float64 {
	// Create a map to hold salaries by age
	salaryByAge := make(map[int][]float64)
	for _, employee := range employees {
		// Calculate employee age
		age := employee.GetAge()
		// Add employee salary to the appropriate age group
		salaryByAge[age] = append(salaryByAge[age], employee.Salary)
	}

	// Calculate the average salary for each age group
	avgSalaryByAgeGroup := make(map[int]float64)
	for age, salaries := range salaryByAge {
		// Determine the age group based on the employee's age
		ageGroup := (age-1)/10 + 1
		// Calculate the total salary for the age group
		totalSalary := 0.0
		for _, salary := range salaries {
			totalSalary += salary
		}
		// Calculate the average salary for the age group
		avgSalaryByAgeGroup[ageGroup] = totalSalary / float64(len(salaries))
	}

	return avgSalaryByAgeGroup
}

// Define the main function of the program
func main() {
	// Open the CSV file containing employee information
	file, err := os.Open("employees.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Create a CSV reader to parse the file
	reader := csv.NewReader(file)
	reader.Comma = ';'
	reader.FieldsPerRecord = 4

	// Read all the records from the CSV file
	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	// Create a slice to hold employee information
	employees := make([]Employee, len(records)-1)

	// Parse each record and create an employee object
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

	// Output employee information to console
	for _, employee := range employees {
		fmt.Printf("%s is %d years old, earns %.2f and is in %s department\n", employee.FullName, employee.GetAge(), employee.Salary, employee.Department)
	}

	// Output average salary groups
	avgSalaryByAgeGroup := averageSalaryByAgeGroup(employees)
	for ageGroup, avgSalary := range avgSalaryByAgeGroup {
		fmt.Printf("Average salary for age group %d is %.2f\n", ageGroup, avgSalary)
	}
}
