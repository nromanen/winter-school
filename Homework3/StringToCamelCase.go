package main

import (
    "fmt"
    "strings"
)

func isTextCap(text string) bool {
    return strings.ToUpper(text[0:1]) == text[0:1]
}

func toCamelCase(text string) string {
    res := ""
    var textInit = strings.Split(text, "-")
    if len(textInit) == 1 {
        textInit = strings.Split(text, "_")
    }

    if !isTextCap(textInit[0]) {
        res += strings.ToLower(textInit[0])
        for i := 1; i < len(textInit); i++ {
            textInit[i] = strings.Title(textInit[i])
        }
    } else {
        res += textInit[0]
        for i := 1; i < len(textInit); i++ {
            textInit[i] = strings.Title(textInit[i])
        }
    }

    res += strings.Join(textInit[1:], "")
    return res
}

func main() {
    fmt.Println(toCamelCase("the-stealth-warrior"))
    fmt.Println(toCamelCase("The_Stealth_Warrior"))
    fmt.Println(toCamelCase("The_stealth_warrior"))
}

