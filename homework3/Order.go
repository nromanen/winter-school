package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Order(sentence string) string {
	words := strings.Fields(sentence)
	for i := 1; i <= len(words); i++ {
		for j := i - 1; j < len(words); j++ {
			if strings.Contains(words[j], strconv.Itoa(i)) {
				words[j], words[i-1] = words[i-1], words[j]
			}
		}
	}
	return strings.Join(words, " ")
}

func main() {
	fmt.Println(Order("is2 Thi1s T4est 3a"))
	fmt.Println(Order("4of Fo1r pe6ople g3ood th5e the2"))
}
