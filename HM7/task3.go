package main

import (
    "encoding/csv"
    "fmt"
    "os"
)

func main() {
    file, err := os.Open("data.csv")
    if err != nil {
        fmt.Println("Помилка відкриття файлу:", err)
        return
    }
    defer file.Close()

    reader := csv.NewReader(file)

    records, err := reader.ReadAll()
    if err != nil {
        fmt.Println("Помилка читання файлу CSV:", err)
        return
    }

    for _, record := range records {
        for _, field := range record {
            fmt.Print(field, " ")
        }
        fmt.Println()
    }
}
