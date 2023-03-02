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

func (e *Employee) GetAge() int {
    today := time.Now()
    age := today.Year() - e.BirthDate.Year()
    if today.Month() < e.BirthDate.Month() ||
        (today.Month() == e.BirthDate.Month() && today.Day() < e.BirthDate.Day()) {
        age--
    }
    return age
}

func main() {
    employees := []Employee{
        Employee{FullName: "John Smith", BirthDate: time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC), Salary: 50000},
        Employee{FullName: "Jane Doe", BirthDate: time.Date(1985, 5, 10, 0, 0, 0, 0, time.UTC), Salary: 60000},
        Employee{FullName: "Bob Johnson", BirthDate: time.Date(2000, 9, 15, 0, 0, 0, 0, time.UTC), Salary: 45000},
        Employee{FullName: "Mary Lee", BirthDate: time.Date(1978, 3, 25, 0, 0, 0, 0, time.UTC), Salary: 70000},
        Employee{FullName: "Bred Hlib", BirthDate: time.Date(2008, 10, 20, 0, 0, 0, 0, time.UTC), Salary: 100000},
        Employee{FullName: "Mary Johnson", BirthDate: time.Date(2010, 3, 23, 0, 0, 0, 0, time.UTC), Salary: 110000},
        Employee{FullName: "Stepan Bandera", BirthDate: time.Date(2012, 6, 26, 0, 0, 0, 0, time.UTC), Salary: 40000},
        Employee{FullName: "John Johnson", BirthDate: time.Date(2010, 4, 24, 0, 0, 0, 0, time.UTC), Salary: 600000},
    }

    ageGroups := make(map[int][]Employee)
    for _, emp := range employees {
        age := emp.GetAge()
        ageGroup := (age-1)/10 + 1 // calculate age group
        ageGroups[ageGroup] = append(ageGroups[ageGroup], emp)
    }

    for ageGroup, emps := range ageGroups {
        avgSalary := 0.0
        if len(emps) > 0 {
            sum := 0.0
            for _, emp := range emps {
                sum += emp.Salary
            }
            avgSalary = sum / float64(len(emps))
        }
        fmt.Printf("Age Group %d:\n", ageGroup)
        for _, emp := range emps {
            fmt.Printf("    %s (Salary: %.2f)\n", emp.FullName, emp.Salary)
        }
        fmt.Printf("    Average Salary: %.2f\n", avgSalary)
    }
}
