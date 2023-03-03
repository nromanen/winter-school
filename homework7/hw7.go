package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
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
	now := time.Now()
	age := now.Year() - e.BirthDate.Year()
	if now.YearDay() < e.BirthDate.YearDay() {
		age--
	}
	return age
}

func main() {

	file, err := os.Open("info.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ';'

	var employees []Employee

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		fullName := record[0]
		birthDate, err := time.Parse("2006-01-02", record[1])
		if err != nil {
			log.Fatal(err)
		}
		salary, err := strconv.ParseFloat(record[2], 64)
		if err != nil {
			log.Fatal(err)
		}

		employee := Employee{fullName, birthDate, salary}
		employees = append(employees, employee)
	}

	ageGroups := make(map[int][]float64)

	for _, employee := range employees {
		age := employee.GetAge()
		ageGroup := age / 10
		ageGroups[ageGroup] = append(ageGroups[ageGroup], employee.Salary)
	}

	for ageGroup, salaries := range ageGroups {
		if len(salaries) == 0 {
			continue
		}
		total := 0.0
		for _, salary := range salaries {
			total += salary
		}
		average := total / float64(len(salaries))
		fmt.Printf("Age group %d: average salary = %.2f\n", ageGroup+1, average)
	}
}
