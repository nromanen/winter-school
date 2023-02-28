package kata

import (
    "strings"
)

func ToCamelCase(s string) string {
    words := strings.Split(s, "-")
    for i, word := range words {
        if strings.Contains(word, "_") {
            words[i] = strings.Title(strings.Replace(word, "_", "", -1))
        }
    }
    return strings.Join(words, "")
}