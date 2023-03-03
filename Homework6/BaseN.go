package main

import (
    "fmt"
)

func convert(a, base int) (string, error) {
    if base < 2 || base > 36 {
        return "", fmt.Errorf("invalid base %d", base)
    }

    const charset = "0123456789abcdefghijklmnopqrstuvwxyz"
    if a < 0 {
        a = -a
    }

    var result string
    for a > 0 {
        remainder := a % base
        result = string(charset[remainder]) + result
        a = a / base
    }

    if result == "" {
        result = "0"
    }

    return result, nil
}

func main() {
    // Виклик функції convert() з аргументами 12 та 16
    result, err := convert(12, 16)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println(result) // Виводить "C"
    }
}
