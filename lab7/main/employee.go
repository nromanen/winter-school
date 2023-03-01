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
	Salary    uint
}

func (e Employee) GetAge() int {
	return time.Now().Year() - e.BirthDate.Year()
}

func main() {
	var employees []Employee

	f, err := os.Open("data.csv")
	if err != nil {
		log.Fatal(err)
	}
	r := csv.NewReader(f)
	// Parse data from csv file
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		date, _ := time.Parse("2006-01-02", record[1])
		salary, _ := strconv.Atoi(record[2])
		employees = append(employees,
			Employee{
				FullName:  record[0],
				BirthDate: date,
				Salary:    uint(salary),
			})
	}
	// Divide into age groups and display average salaries for each
	fmt.Println("Average salaries")
	var sum, c uint
	ageGroup := 10
	for i := 0; (i+1)*ageGroup < 80; i++ {
		sum = 0
		c = 0
		for _, v := range employees {
			if i*ageGroup < v.GetAge() && v.GetAge() < (i+1)*ageGroup {
				fmt.Println(v)
				sum += v.Salary
				c++
			}
		}
		if c > 0 {
			fmt.Printf("%v - %v: $%v\n", i*ageGroup, (i+1)*ageGroup, float32(sum/c))
		}
	}
	f.Close()
}
