package main

import (
    "fmt"
)

type Employee struct {
    FullName string
    Age      int
    Salary   float64
}

func main() {
    employees := []Employee{
        {"John Doe", 25, 35000.0},
        {"Jane Doe", 35, 50000.0},
        {"Bob Smith", 42, 65000.0},
        {"Mary Johnson", 18, 25000.0},
        {"Tom Wilson", 28, 40000.0},
        {"Alice Brown", 23, 32000.0},
    }

    // встановлюємо розміри вікових груп
    groupSizes := make(map[int]int)
    for i := 1; i <= 10; i++ {
        groupSizes[i] = 0
    }

    // знаходимо зарплати для кожної вікової групи
    groupSalaries := make(map[int][]float64)
    for _, employee := range employees {
        ageGroup := (employee.Age-1)/10 + 1
        groupSizes[ageGroup]++
        groupSalaries[ageGroup] = append(groupSalaries[ageGroup], employee.Salary)
    }

    // виводимо середню зарплату для кожної вікової групи
    for i := 1; i <= 10; i++ {
        if groupSizes[i] == 0 {
            continue
        }
        averageSalary := sum(groupSalaries[i]) / float64(groupSizes[i])
        fmt.Printf("Average salary for group %d: %.2f\n", i, averageSalary)
    }
}

func sum(values []float64) float64 {
    result := 0.0
    for _, value := range values {
        result += value
    }
    return result
}
