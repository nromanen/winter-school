package main

import (
    "fmt"
    "time"
)

type Employee struct {
    FullName  string
    BirthDate time.Time
    Salary    int
}

func (e Employee) GetAge() int {
    now := time.Now()
    age := now.Year() - e.BirthDate.Year()
    if now.Month() < e.BirthDate.Month() || (now.Month() == e.BirthDate.Month() && now.Day() < e.BirthDate.Day()) {
        age--
    }
    return age
}

func main() {
    birthDate, _ := time.Parse("2006-01-02", "1990-05-10")
    employee := Employee{FullName: "John Doe", BirthDate: birthDate, Salary: 50000}

    fmt.Println("Employee Name:", employee.FullName)
    fmt.Println("Employee Age:", employee.GetAge())
    fmt.Println("Employee Salary:", employee.Salary)
}
