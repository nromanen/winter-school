package main

import (
    "fmt"
    "strings"
)

// Повертає число у заданій системі числення
func convert(a, base int) (string, error) {
    if base < 2 || base > 36 {
        return "", fmt.Errorf("невірна система числення: %d", base)
    }

    const charset = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
    var result strings.Builder

    // Розраховуємо числа у зворотному порядку
    for a > 0 {
        result.WriteByte(charset[a%base])
        a /= base
    }

    // Перевертаємо рядок
    reversed := result.String()
    length := len(reversed)
    runes := make([]rune, length)
    for i, r := range reversed {
        runes[length-1-i] = r
    }

    return string(runes), nil
}

func main() {
    n := 12
    base := 16
    result, err := convert(n, base)
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Printf("%d у системі числення %d = %s\n", n, base, result)
    }
}