package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"sort"
	"time"
)

type Employee struct {
	FullName  string
	BirthDate time.Time
	Salary    float64
}

func (e *Employee) GetAge() int {
	now := time.Now()
	diff := now.Sub(e.BirthDate)
	age := int(diff.Hours() / 24 / 365)
	return age
}

func main() {
	file, err := os.Open("Data.csv")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ';'
	var employees []Employee

	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		birthdate, _ := time.Parse("02.01.2006", row[1])
		salary := 0.0
		fmt.Sscanf(row[2], "%f", &salary)

		employee := Employee{
			FullName:  row[0],
			BirthDate: birthdate,
			Salary:    salary,
		}
		employees = append(employees, employee)
	}

	ageGroups := make(map[int][]float64)
	for _, employee := range employees {
		age := employee.GetAge()
		group := age / 10
		if group > 3 {
			continue // ignore age group 31-40 and above
		}
		ageGroups[group] = append(ageGroups[group], employee.Salary)
	}

	var keys []int
	for k := range ageGroups {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, group := range keys {
		salaries := ageGroups[group]
		total := 0.0
		for _, salary := range salaries {
			total += salary
		}
		avg := int(total / float64(len(salaries)))
		fmt.Printf("Average salary for age group %d-%d is %d\n", group*10+1, (group+1)*10, avg)
	}
}
