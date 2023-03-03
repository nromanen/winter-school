package main

import (
    "encoding/csv"
    "fmt"
    "os"
    "strconv"
)

type Employee struct {
    FullName  string
    BirthDate string
    Salary    float64
}

func main() {
    file, err := os.Open("employees.csv")
    if err != nil {
        fmt.Println("Помилка відкриття файлу:", err)
        return
    }
    defer file.Close()

    reader := csv.NewReader(file)

    records, err := reader.ReadAll()
    if err != nil {
        fmt.Println("Помилка зчитування файлу:", err)
        return
    }

    ageGroups := make(map[int][]float64)

    for _, record := range records {
        salary, _ := strconv.ParseFloat(record[2], 64) // Перетворюємо зарплату у числовий тип
        birthYear, _ := strconv.Atoi(record[1][:4])    // Отримуємо рік народження працівника
        ageGroup := (2023 - birthYear) / 10            // Обчислюємо вікову групу працівника
        ageGroups[ageGroup] = append(ageGroups[ageGroup], salary) // Додаємо
